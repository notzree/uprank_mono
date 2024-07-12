package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		const (
			env              = "dev"
			application_name = "uprank"
		)
		// Read the .env file
		envFileContent, err := os.ReadFile("../../../.env")
		if err != nil {
			return fmt.Errorf("failed to read .env file: %w", err)
		}
		// Convert .env content to JSON
		envMap := make(map[string]string)
		lines := strings.Split(string(envFileContent), "\n")
		for _, line := range lines {
			if line != "" && !strings.HasPrefix(line, "#") {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					envMap[parts[0]] = parts[1]
				}
			}
		}

		envJSON, err := json.Marshal(envMap)
		if err != nil {
			return fmt.Errorf("failed to marshal env map to JSON: %w", err)
		}

		// Create a secret in AWS Secrets Manager
		secret, err := secretsmanager.NewSecret(ctx, "backend-secrets", &secretsmanager.SecretArgs{})
		if err != nil {
			return fmt.Errorf("failed to create secret: %w", err)
		}

		// Store the contents of the .env file in the secret
		_, err = secretsmanager.NewSecretVersion(ctx, CreateResourceName(env, application_name, "backend-secrets"), &secretsmanager.SecretVersionArgs{
			SecretId:     secret.ID(),
			SecretString: pulumi.String(envJSON),
		})
		if err != nil {
			return fmt.Errorf("failed to create secret version: %w", err)
		}

		ctx.Export("secretArn", secret.Arn)
		return nil
	})
}

func CreateResourceName(env string, application_name string, resource string) string {
	return fmt.Sprintf("%s-%s-%s", application_name, resource, env)
}
