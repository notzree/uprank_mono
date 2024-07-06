package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	const (
		env              = "local"
		application_name = "uprank"
	)
	pulumi.Run(func(ctx *pulumi.Context) error {

		ranking_queue, err := sqs.NewQueue(ctx, CreateResourceName(env, application_name, "ranking_queue"), &sqs.QueueArgs{
			FifoQueue:                 pulumi.Bool(true),
			ContentBasedDeduplication: pulumi.Bool(true),
			SqsManagedSseEnabled:      pulumi.Bool(true),
			VisibilityTimeoutSeconds:  pulumi.Int(200),
		})
		if err != nil {
			return err
		}
		notification_queue, err := sqs.NewQueue(ctx, CreateResourceName(env, application_name, "notification_queue"), &sqs.QueueArgs{
			FifoQueue:                 pulumi.Bool(true),
			ContentBasedDeduplication: pulumi.Bool(true),
			SqsManagedSseEnabled:      pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		ctx.Export("ranking_queue_url", ranking_queue.Url)
		ctx.Export("ranking_queue_arn", ranking_queue.Arn)
		ctx.Export("notification_queue_url", notification_queue.Url)
		return nil
	})
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
