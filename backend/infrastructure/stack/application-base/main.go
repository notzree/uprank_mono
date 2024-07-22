package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	ecrx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

//Required ECS Service + other resources such as ALB, Security Group, Service Discovery, etc.

func main() {
	pulumi.Run(pulumi.RunFunc(func(ctx *pulumi.Context) error {
		const (
			application_name = "uprank"
		)
		stack := ctx.Stack()

		networking_repository, err := pulumi.NewStackReference(ctx, fmt.Sprintf("notzree/networking/%s", stack), nil)
		if err != nil {
			return err
		}
		vpc_id := networking_repository.GetOutput(pulumi.String("vpc_id"))

		// An ECS cluster to deploy into
		cluster, err := ecs.NewCluster(ctx, CreateResourceName(stack, application_name, "cluster"), nil)
		if err != nil {
			return err
		}
		// ECR repository for all images
		repo, err := ecrx.NewRepository(ctx, CreateResourceName(stack, application_name, "uprank-repo"), &ecrx.RepositoryArgs{
			ForceDelete: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// DNS discovery service
		private_dns_namespace, err := servicediscovery.NewPrivateDnsNamespace(ctx, CreateResourceName(stack, application_name, "dns_namespace"), &servicediscovery.PrivateDnsNamespaceArgs{
			Name:        pulumi.String("dev.uprank.ca"),
			Description: pulumi.String("Development namespace for Uprank"),
			Vpc:         pulumi.StringOutput(vpc_id),
		})
		if err != nil {
			return err
		}

		//todo: register services to the service discovery
		ctx.Export("ecr_url", repo.Url)
		ctx.Export("ecs_cluster_arn", cluster.Arn)
		ctx.Export("private_dns_namespace_id", private_dns_namespace.ID())
		return nil
	}))
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
