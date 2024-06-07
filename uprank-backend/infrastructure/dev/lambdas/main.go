package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		queue_stack, err := pulumi.NewStackReference(ctx, "notzree/queues/dev", nil)
		if err != nil {
			return err
		}
		queueArn := queue_stack.GetOutput(pulumi.String("scraper_queue_arn"))

		// Create IAM role and policy for Lambda
		role, err := iam.NewRole(ctx, "lambdaRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Action": "sts:AssumeRole",
						"Principal": {
							"Service": "lambda.amazonaws.com"
						},
						"Effect": "Allow"
					}
				]
			}`),
		})
		if err != nil {
			return err
		}

		_, err = iam.NewRolePolicy(ctx, "lambdaPolicy", &iam.RolePolicyArgs{
			Role: role.Name,
			Policy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Action": [
							"logs:CreateLogGroup",
							"logs:CreateLogStream",
							"logs:PutLogEvents",
							"sqs:ReceiveMessage",
							"sqs:DeleteMessage",
							"sqs:GetQueueAttributes"
						],
						"Resource": "*"
					}
				]
			}`),
		})
		if err != nil {
			return err
		}

		bucket, err := s3.NewBucket(ctx, "scraper-deps-layer-bucket", nil)
		if err != nil {
			return err
		}

		layer_object, err := s3.NewBucketObject(ctx, "lambda-layer-zip", &s3.BucketObjectArgs{
			Bucket: bucket.ID(),
			Source: pulumi.NewFileAsset("../../../scraper/layers/dependencies.zip"),
		})
		if err != nil {
			return err
		}
		lambda_layer, err := lambda.NewLayerVersion(ctx, "myLambdaLayer", &lambda.LayerVersionArgs{
			CompatibleRuntimes: pulumi.StringArray{
				pulumi.String("nodejs20.x"), // Specify the runtimes your layer is compatible with
			},
			CompatibleArchitectures: pulumi.StringArray{
				pulumi.String("x86_64"),
			},
			S3Bucket:    bucket.ID(),
			S3Key:       layer_object.Key,
			LayerName:   pulumi.String("scraper-dependency-layer"),
			Description: pulumi.String("Browser binaries and other dependencies"),
		})
		if err != nil {
			return err
		}

		// Create Lambda function
		fn, err := lambda.NewFunction(ctx, "uprank-scraper-function", &lambda.FunctionArgs{
			Code:       pulumi.NewFileArchive("../../../scraper/dist/index.zip"),
			Handler:    pulumi.String("index.handler"),
			Role:       role.Arn,
			MemorySize: pulumi.Int(1024),
			Timeout:    pulumi.Int(300),
			Runtime:    pulumi.String("nodejs20.x"),
			Layers: pulumi.StringArray{
				lambda_layer.Arn,
			},
		})
		if err != nil {
			return err
		}

		// Create Event Source Mapping
		_, err = lambda.NewEventSourceMapping(ctx, "eventSourceMapping", &lambda.EventSourceMappingArgs{
			EventSourceArn: pulumi.StringOutput(queueArn),
			FunctionName:   fn.Name,
			Enabled:        pulumi.Bool(true),
			BatchSize:      pulumi.Int(10),
		})
		if err != nil {
			return err
		}

		ctx.Export("lambdaFunctionName", fn.Name)
		return nil
	})
}
