package main

import (
	"fmt"

	"github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		const (
			application_name = "uprank"
		)
		stack := ctx.Stack()
		// Allocate a new VPC with the default settings.
		vpc, err := ec2.NewVpc(ctx, CreateResourceName(stack, application_name, "vpc"), &ec2.VpcArgs{
			EnableDnsSupport:   pulumi.Bool(true),
			EnableDnsHostnames: pulumi.Bool(true),
			NatGateways: &ec2.NatGatewayConfigurationArgs{
				Strategy: ec2.NatGatewayStrategyOnePerAz,
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("vpc_id", vpc.VpcId)
		ctx.Export("private_subnet_ids", vpc.PrivateSubnetIds)
		ctx.Export("public_subnet_ids", vpc.PublicSubnetIds)
		return nil
	})
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
