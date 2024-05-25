package main

import (
	"fmt"

	awsec2 "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	ecrx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecr"
	ecsx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecs"
	lbx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// File for local development. Will only contain the queues that are required.

func main() {
	const (
		env              = "local"
		application_name = "uprank"
	)

	pulumi.Run(pulumi.RunFunc(func(ctx *pulumi.Context) error {
		//todo: fix config stuff, also change the implementation of the services
		//todo: add the scraper and main backend services
		// probably don't need nginx? idk yet
		container_repository, err := pulumi.NewStackReference(ctx, "notzree/container-repository/dev", nil)
		if err != nil {
			return err
		}
		container_repository_url := container_repository.GetOutput(pulumi.String("repository_url"))

		networking_repository, err := pulumi.NewStackReference(ctx, "notzree/networking/dev", nil)
		if err != nil {
			return err
		}
		vpc_id := networking_repository.GetOutput(pulumi.String("vpc_id"))
		// private_subnet_ids := networking_repository.GetOutput(pulumi.String("private_subnet_ids"))
		public_subnet_ids := networking_repository.GetOutput(pulumi.String("public_subnet_ids"))

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

		securityGroup, err := awsec2.NewSecurityGroup(ctx, "security_group", &awsec2.SecurityGroupArgs{
			VpcId: pulumi.StringOutput(vpc_id),
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

		// An ECS cluster to deploy into
		cluster, err := ecs.NewCluster(ctx, CreateResourceName(env, application_name, "cluster"), nil)
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

		// Build and publish our application's container image from ./app to the ECR repository
		image, err := ecrx.NewImage(ctx, "image", &ecrx.ImageArgs{
			RepositoryUrl: pulumi.StringOutput(container_repository_url),
			Context:       pulumi.String("../../../main-backend/"),
			Dockerfile:    pulumi.String("../../../main-backend/Dockerfile"),
			Platform:      pulumi.String("linux/amd64"),
		})
		if err != nil {
			return err
		}

		// Deploy an ECS Service on Fargate to host the application container
		// https://stackoverflow.com/questions/61265108/aws-ecs-fargate-resourceinitializationerror-unable-to-pull-secrets-or-registry
		_, err = ecsx.NewFargateService(ctx, CreateResourceName(env, application_name, "fargate_service"), &ecsx.FargateServiceArgs{
			Cluster: cluster.Arn,
			NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
				AssignPublicIp: pulumi.Bool(true),
				Subnets:        pulumi.StringArrayOutput(public_subnet_ids),
				SecurityGroups: pulumi.StringArray{
					securityGroup.ID(),
				},
			},
			// AssignPublicIp: pulumi.Bool(true),
			TaskDefinitionArgs: &ecsx.FargateServiceTaskDefinitionArgs{

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
				},
			},
		})
		if err != nil {
			return err
		}
		// The URL at which the container's HTTP endpoint will be available
		ctx.Export("url", pulumi.Sprintf("http://%s", loadbalancer.LoadBalancer.DnsName()))
		return nil
	}))
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
