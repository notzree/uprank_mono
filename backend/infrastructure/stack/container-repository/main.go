package main

import (
	"fmt"

	ecrx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	const (
		env              = "dev"
		application_name = "uprank"
	)
	pulumi.Run(func(ctx *pulumi.Context) error {
		repo, err := ecrx.NewRepository(ctx, CreateResourceName(env, application_name, "repo"), &ecrx.RepositoryArgs{
			ForceDelete: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("repository_url", repo.Url)
		return nil
	})
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
