package main

import (
	"fmt"
	"time"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	awsec2 "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/awsx"
	ecrx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecr"
	ecsx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecs"
	lbx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stack := ctx.Stack()
		const (
			application_name = "uprank"
		)
		// <-- Stack references -->
		application_base, err := pulumi.NewStackReference(ctx, fmt.Sprintf("notzree/application-base/%s", stack), nil)
		if err != nil {
			return err
		}
		ecr_url := application_base.GetOutput(pulumi.String("ecr_url"))
		private_dns_namespace_id := application_base.GetOutput(pulumi.String("private_dns_namespace_id"))
		cluster_arn := application_base.GetOutput(pulumi.String("ecs_cluster_arn"))
		networking_repository, err := pulumi.NewStackReference(ctx, fmt.Sprintf("notzree/networking/%s", stack), nil)
		if err != nil {
			return err
		}
		private_subnet_ids := networking_repository.GetOutput(pulumi.String("private_subnet_ids"))
		public_subnet_ids := networking_repository.GetOutput(pulumi.String("public_subnet_ids"))
		vpc_id := networking_repository.GetOutput(pulumi.String("vpc_id"))

		secret_repository, err := pulumi.NewStackReference(ctx, fmt.Sprintf("notzree/secrets/%s", stack), nil)
		if err != nil {
			return err
		}
		secret_arn := secret_repository.GetOutput(pulumi.String("secretArn"))
		// <-- End Stack references -->

		cfg := config.New(ctx, "")
		containerPort := 80
		if param := cfg.GetInt("containerPort"); param != 0 {
			containerPort = param
		}
		cpu := 512
		if param := cfg.GetInt("cpu"); param != 0 {
			cpu = param
		}
		memory := 1024
		if param := cfg.GetInt("memory"); param != 0 {
			memory = param
		}

		image, err := ecrx.NewImage(ctx, CreateImageName(stack, application_name, "main-backend"), &ecrx.ImageArgs{
			RepositoryUrl: pulumi.StringOutput(ecr_url),
			Context:       pulumi.String("../../../main-backend"),
			Dockerfile:    pulumi.String("../../../main-backend/Dockerfile"),
			Platform:      pulumi.String("linux/amd64"),
		})
		if err != nil {
			return err
		}

		// Create IAM Role for ECS Task Execution
		taskRole, err := iam.NewRole(ctx, "main-backend-ecs-task-role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
						"Version": "2012-10-17",
						"Statement": [
							{
								"Effect": "Allow",
								"Principal": {
									"Service": "ecs-tasks.amazonaws.com"
								},
								"Action": "sts:AssumeRole"
							}
						]
					}`),
		})
		if err != nil {
			return err
		}
		// Create role policy to access SQS
		_, err = iam.NewRolePolicy(ctx, "ecsSQSAccess", &iam.RolePolicyArgs{
			Role: taskRole.Name,
			Policy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Action": [
							"sqs:receivemessage",
							"sqs:DeleteMessage",
							"sqs:ListQueues",
							"sqs:GetQueueAttributes",
							"sqs:CancelMessageMoveTask",
							"sqs:SendMessage"
						],
						"Resource": "*"
					}
				]
			}`),
		})
		if err != nil {
			return err
		}

		// Create IAM Role for ECS Task Execution
		executionRole, err := iam.NewRole(ctx, "main-backend-ecs-task-execution-role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
						"Version": "2012-10-17",
						"Statement": [
							{
								"Effect": "Allow",
								"Principal": {
									"Service": "ecs-tasks.amazonaws.com"
								},
								"Action": "sts:AssumeRole"
							}
						]
					}`),
		})
		if err != nil {
			return err
		}

		// Attach the policy to the IAM Role
		_, err = iam.NewRolePolicyAttachment(ctx, "ecsTaskExecutionRolePolicy", &iam.RolePolicyAttachmentArgs{
			Role:      executionRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
		})
		if err != nil {
			return err
		}

		// Attach the custom Secrets Manager policy to the IAM Role
		_, err = iam.NewRolePolicy(ctx, "ecsTaskSecretsManagerPolicy", &iam.RolePolicyArgs{
			Role: executionRole.Name,
			Policy: pulumi.String(`{
						"Version": "2012-10-17",
						"Statement": [
							{
								"Effect": "Allow",
								"Action": [
									"secretsmanager:GetSecretValue"
								],
								"Resource": "*"
							}
						]
					}`),
		})
		if err != nil {
			return err
		}

		_, err = iam.NewRolePolicy(ctx, "ecsTaskCloudWatchPolicy", &iam.RolePolicyArgs{
			Role: executionRole.Name,
			Policy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Action": [
							"logs:CreateLogStream",
							"logs:PutLogEvents"
						],
						"Resource": "arn:aws:logs:*:*:*"
					}
				]
			}`),
		})
		if err != nil {
			return err
		}

		// Create the CloudWatch Log Group
		logGroup, err := cloudwatch.NewLogGroup(ctx, "uprank-main-backend-log-group", &cloudwatch.LogGroupArgs{
			Name: pulumi.String("/ecs/main-backend-log-group"),
		})
		if err != nil {
			return fmt.Errorf("failed to create log group: %w", err)
		}

		securityGroup, err := awsec2.NewSecurityGroup(ctx, "main-backend-security-group", &awsec2.SecurityGroupArgs{
			VpcId: pulumi.StringOutput(vpc_id),
			Ingress: awsec2.SecurityGroupIngressArray{
				&awsec2.SecurityGroupIngressArgs{
					FromPort:   pulumi.Int(80), //custom http
					ToPort:     pulumi.Int(80),
					Protocol:   pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
				&awsec2.SecurityGroupIngressArgs{ //default https
					FromPort:   pulumi.Int(443),
					ToPort:     pulumi.Int(443),
					Protocol:   pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
			Egress: awsec2.SecurityGroupEgressArray{
				&awsec2.SecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
					Ipv6CidrBlocks: pulumi.StringArray{
						pulumi.String("::/0"),
					},
				},
			},
		})

		if err != nil {
			return err
		}

		// An ALB to serve the container endpoint to the internet
		loadbalancer, err := lbx.NewApplicationLoadBalancer(ctx, CreateResourceName(stack, application_name, "alb"), &lbx.ApplicationLoadBalancerArgs{
			SubnetIds: pulumi.StringArrayOutput(public_subnet_ids),
			// DefaultTargetGroupPort: pulumi.Int(80), //default http port
			DefaultTargetGroup: &lbx.TargetGroupArgs{
				Name:     pulumi.String("main-backend-target-group"),
				Port:     pulumi.Int(80), //forward to default http port
				Protocol: pulumi.String("HTTP"),
				HealthCheck: lb.TargetGroupHealthCheckPtrInput(lb.TargetGroupHealthCheckArgs{
					Port: pulumi.String("traffic-port"),
					Path: pulumi.String("/healthz"),
				}),
			},
		})
		if err != nil {
			return err
		}

		main_backend_service_discovery, err := servicediscovery.NewService(ctx, CreateResourceName(stack, application_name, "main-backend"), &servicediscovery.ServiceArgs{
			Name:         pulumi.String("main-backend-servicediscovery"),
			ForceDestroy: pulumi.Bool(true),
			DnsConfig: &servicediscovery.ServiceDnsConfigArgs{
				NamespaceId: pulumi.StringOutput(private_dns_namespace_id),
				DnsRecords: servicediscovery.ServiceDnsConfigDnsRecordArray{
					&servicediscovery.ServiceDnsConfigDnsRecordArgs{
						Ttl: pulumi.Int(30),

						Type: pulumi.String("A"),
					},
				},
				RoutingPolicy: pulumi.String("MULTIVALUE"),
			},
			HealthCheckCustomConfig: &servicediscovery.ServiceHealthCheckCustomConfigArgs{
				FailureThreshold: pulumi.Int(3),
			},
		})
		if err != nil {
			return err
		}

		_, err = ecsx.NewFargateService(ctx, CreateResourceName(stack, application_name, "main-backend-service"), &ecsx.FargateServiceArgs{
			Cluster: pulumi.StringOutput(cluster_arn),
			NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
				AssignPublicIp: pulumi.Bool(false),
				Subnets:        pulumi.StringArrayOutput(private_subnet_ids),
				SecurityGroups: pulumi.StringArray{
					securityGroup.ID(),
				},
			},
			HealthCheckGracePeriodSeconds: pulumi.Int(30),
			ServiceRegistries: ecs.ServiceServiceRegistriesPtrInput(
				&ecs.ServiceServiceRegistriesArgs{
					RegistryArn: main_backend_service_discovery.Arn,
				}),
			//todo: transition from service discovery to service connect
			// ServiceConnectConfiguration: ecs.ServiceServiceConnectConfigurationPtrInput(ecs.ServiceServiceConnectConfigurationArgs{
			// 	Enabled:   pulumi.Bool(true),
			// 	Namespace: pulumi.String("dev.uprank.ca"),
			// 	Services:  ecs.ServiceServiceConnectConfigurationServiceArray{},
			// }),
			TaskDefinitionArgs: &ecsx.FargateServiceTaskDefinitionArgs{
				LogGroup: &awsx.DefaultLogGroupArgs{
					Existing: &awsx.ExistingLogGroupArgs{
						Arn: logGroup.Arn,
					},
				},
				TaskRole: &awsx.DefaultRoleWithPolicyArgs{
					RoleArn: taskRole.Arn,
				},
				ExecutionRole: &awsx.DefaultRoleWithPolicyArgs{
					RoleArn: executionRole.Arn,
				},
				Container: &ecsx.TaskDefinitionContainerDefinitionArgs{
					Image:     image.ImageUri,
					Cpu:       pulumi.Int(cpu),
					Memory:    pulumi.Int(memory),
					Essential: pulumi.Bool(true),
					PortMappings: ecsx.TaskDefinitionPortMappingArray{
						&ecsx.TaskDefinitionPortMappingArgs{
							ContainerPort: pulumi.Int(containerPort),
							TargetGroup:   loadbalancer.DefaultTargetGroup,
						},
					},
					HealthCheck: ecsx.TaskDefinitionHealthCheckPtrInput(&ecsx.TaskDefinitionHealthCheckArgs{
						StartPeriod: pulumi.Int(60),
						Command:     pulumi.StringArrayInput(pulumi.ToStringArray([]string{"CMD-SHELL", "curl -f http://localhost:80/healthz || exit 1"})),
					}),
					Secrets: ecsx.TaskDefinitionSecretArray{
						&ecsx.TaskDefinitionSecretArgs{
							Name:      pulumi.String("MAIN_BACKEND_SECRETS"),
							ValueFrom: pulumi.StringOutput(secret_arn),
						},
					},
					Environment: ecsx.TaskDefinitionKeyValuePairArrayInput(ecsx.TaskDefinitionKeyValuePairArray{
						&ecsx.TaskDefinitionKeyValuePairArgs{
							Name:  pulumi.String("ENV"),
							Value: pulumi.String("dev"),
						},
					}),
				},
			},
		},
			pulumi.DependsOn([]pulumi.Resource{loadbalancer}),
		)
		if err != nil {
			return err
		}
		ctx.Export("main-backend-image-uri", image.ImageUri)
		ctx.Export("main-backend-alb-url", pulumi.Sprintf("http://%s", loadbalancer.LoadBalancer.DnsName()))
		return nil
	})
}
func CreateImageName(env string, application_name string, resource string) string {
	currentDate := time.Now().Format("YYYY-MM-DD") // Format the date as YYYY-MM-DD
	return fmt.Sprintf("%s-%s-%s-%s", currentDate, application_name, resource, env)
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
