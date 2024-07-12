package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

//Required ECS Service + other resources such as ALB, Security Group, Service Discovery, etc.

func main() {
	const (
		env              = "local"
		application_name = "uprank"
	)

	pulumi.Run(pulumi.RunFunc(func(ctx *pulumi.Context) error {
		//todo: fix config stuff, also change the implementation of the services
		//todo: add the scraper and main backend services
		// probably don't need nginx? idk yet

		networking_repository, err := pulumi.NewStackReference(ctx, "notzree/networking/dev", nil)
		if err != nil {
			return err
		}
		vpc_id := networking_repository.GetOutput(pulumi.String("vpc_id"))

		// An ECS cluster to deploy into
		cluster, err := ecs.NewCluster(ctx, CreateResourceName(env, application_name, "cluster"), nil)
		if err != nil {
			return err
		}

		// DNS discovery service
		private_dns_namespace, err := servicediscovery.NewPrivateDnsNamespace(ctx, CreateResourceName(env, application_name, "dns_namespace"), &servicediscovery.PrivateDnsNamespaceArgs{
			Name:        pulumi.String("dev.uprank.ca"),
			Description: pulumi.String("Development namespace for Uprank"),
			Vpc:         pulumi.StringOutput(vpc_id),
		})
		if err != nil {
			return err
		}
		_, err = servicediscovery.NewService(ctx, CreateResourceName(env, application_name, "service-discovery"), &servicediscovery.ServiceArgs{
			Name: pulumi.String("service-discovery"),
			DnsConfig: &servicediscovery.ServiceDnsConfigArgs{
				NamespaceId: private_dns_namespace.ID(),
				DnsRecords: servicediscovery.ServiceDnsConfigDnsRecordArray{
					&servicediscovery.ServiceDnsConfigDnsRecordArgs{
						Ttl:  pulumi.Int(30),
						Type: pulumi.String("A"),
					},
				},
				RoutingPolicy: pulumi.String("MULTIVALUE"),
			},
			HealthCheckCustomConfig: &servicediscovery.ServiceHealthCheckCustomConfigArgs{
				FailureThreshold: pulumi.Int(1),
			},
		})
		if err != nil {
			return err
		}

		//todo: register services to the service discovery

		ctx.Export("ecs_cluster_arn", cluster.Arn)
		return nil
	}))
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
