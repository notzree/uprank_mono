package main

import (
	"fmt"
	"time"

	ecrx "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	//this can be used to push new images to the repository
	pulumi.Run(func(ctx *pulumi.Context) error {
		const (
			env              = "dev"
			application_name = "uprank"
		)

		container_repository, err := pulumi.NewStackReference(ctx, "notzree/container-repository/dev", nil)
		if err != nil {
			return err
		}
		container_repository_url := container_repository.GetOutput(pulumi.String("repository_url"))

		image, err := ecrx.NewImage(ctx, CreateImageName(env, application_name, "main-backend"), &ecrx.ImageArgs{
			RepositoryUrl: pulumi.StringOutput(container_repository_url),
			Context:       pulumi.String("../../../queue-handler"),
			Dockerfile:    pulumi.String("../../../queue-handler/Dockerfile.dev"),
			Platform:      pulumi.String("linux/amd64"),
		})
		if err != nil {
			return err
		}
		ctx.Export("image_uri", image.ImageUri)
		return nil
	})
}
func CreateImageName(env string, application_name string, resource string) string {
	currentDate := time.Now().Format("YYYY-MM-DD") // Format the date as YYYY-MM-DD
	return fmt.Sprintf("%s-%s-%s-%s", currentDate, application_name, resource, env)
}
