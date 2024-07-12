package main

import (
	"fmt"
	"time"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	awsec2 "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/awsx"
	ecrx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecr"
	ecsx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecs"
	lbx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	//this can be used to push new images to the repository
	pulumi.Run(func(ctx *pulumi.Context) error {
		const (
			env              = "dev"
			application_name = "uprank"
		)

		application_base, err := pulumi.NewStackReference(ctx, "notzree/application_base/dev", nil)
		if err != nil {
			return err
		}
		ecr_url := application_base.GetOutput(pulumi.String("ecr_url"))

		container_service, err := pulumi.NewStackReference(ctx, "notzree/container-service/dev", nil)
		if err != nil {
			return err
		}
		cluster_arn := container_service.GetOutput(pulumi.String("ecs_cluster_arn"))
		networking_repository, err := pulumi.NewStackReference(ctx, "notzree/networking/dev", nil)
		if err != nil {
			return err
		}
		// private_subnet_ids := networking_repository.GetOutput(pulumi.String("private_subnet_ids"))
		public_subnet_ids := networking_repository.GetOutput(pulumi.String("public_subnet_ids"))
		vpc_id := networking_repository.GetOutput(pulumi.String("vpc_id"))

		secret_repository, err := pulumi.NewStackReference(ctx, "notzree/secrets/dev", nil)
		if err != nil {
			return err
		}
		secret_arn := secret_repository.GetOutput(pulumi.String("secretArn"))

		cfg := config.New(ctx, "")
		containerPort := 8080
		if param := cfg.GetInt("containerPort"); param != 0 {
			containerPort = param
		}
		cpu := 512
		if param := cfg.GetInt("cpu"); param != 0 {
			cpu = param
		}
		memory := 128
		if param := cfg.GetInt("memory"); param != 0 {
			memory = param
		}

		image, err := ecrx.NewImage(ctx, CreateImageName(env, application_name, "main-backend"), &ecrx.ImageArgs{
			RepositoryUrl: pulumi.StringOutput(ecr_url),
			Context:       pulumi.String("../../../main-backend"),
			Dockerfile:    pulumi.String("../../../main-backend/Dockerfile.dev"),
			Platform:      pulumi.String("linux/amd64"),
		})
		if err != nil {
			return err
		}

		// Create IAM Role for ECS Task Execution
		taskRole, err := iam.NewRole(ctx, "ecsTaskExecutionRole", &iam.RoleArgs{
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
			Role:      taskRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
		})
		if err != nil {
			return err
		}

		// Attach the custom Secrets Manager policy to the IAM Role
		_, err = iam.NewRolePolicy(ctx, "ecsTaskSecretsManagerPolicy", &iam.RolePolicyArgs{
			Role: taskRole.Name,
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
			Role: taskRole.Name,
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
			Name: pulumi.String("/ecs/myLogGroup"),
		})
		if err != nil {
			return fmt.Errorf("failed to create log group: %w", err)
		}

		securityGroup, err := awsec2.NewSecurityGroup(ctx, "main-backend-security-group", &awsec2.SecurityGroupArgs{
			VpcId: pulumi.StringOutput(vpc_id),
			Ingress: awsec2.SecurityGroupIngressArray{
				&awsec2.SecurityGroupIngressArgs{
					FromPort:   pulumi.Int(80),
					ToPort:     pulumi.Int(80),
					Protocol:   pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
				&awsec2.SecurityGroupIngressArgs{
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
		loadbalancer, err := lbx.NewApplicationLoadBalancer(ctx, CreateResourceName(env, application_name, "alb"), &lbx.ApplicationLoadBalancerArgs{
			SubnetIds:              pulumi.StringArrayOutput(public_subnet_ids),
			DefaultTargetGroupPort: pulumi.Int(containerPort),
		})
		if err != nil {
			return err
		}

		_, err = ecsx.NewFargateService(ctx, CreateResourceName(env, application_name, "fargate_service"), &ecsx.FargateServiceArgs{
			Cluster: pulumi.StringOutput(cluster_arn),
			NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
				AssignPublicIp: pulumi.Bool(true),
				Subnets:        pulumi.StringArrayOutput(public_subnet_ids),
				SecurityGroups: pulumi.StringArray{
					securityGroup.ID(),
				},
			},
			// AssignPublicIp: pulumi.Bool(true),
			TaskDefinitionArgs: &ecsx.FargateServiceTaskDefinitionArgs{
				LogGroup: &awsx.DefaultLogGroupArgs{
					Existing: &awsx.ExistingLogGroupArgs{
						Arn: logGroup.Arn,
					},
				},
				ExecutionRole: &awsx.DefaultRoleWithPolicyArgs{
					RoleArn: taskRole.Arn,
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
					Secrets: ecsx.TaskDefinitionSecretArray{
						&ecsx.TaskDefinitionSecretArgs{
							Name:      pulumi.String("SECRET"),
							ValueFrom: pulumi.StringOutput(secret_arn),
						},
					},
				},
			},
		})
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
