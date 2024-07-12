package main

import (
	"github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Allocate a new VPC with the default settings.
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			EnableDnsSupport:   pulumi.Bool(true),
			EnableDnsHostnames: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Export a few properties to make them easy to use.
		ctx.Export("vpc_id", vpc.VpcId)
		ctx.Export("private_subnet_ids", vpc.PrivateSubnetIds)
		ctx.Export("public_subnet_ids", vpc.PublicSubnetIds)
		return nil
	})
}
