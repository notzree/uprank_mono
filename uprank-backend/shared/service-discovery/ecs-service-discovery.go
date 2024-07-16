package servicediscovery

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

type ECSServiceDiscoveryClient struct {
	namespace string
	svc       *servicediscovery.Client
}

func NewECSServiceDiscoveryClient(namespace string) (*ECSServiceDiscoveryClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		return nil, err
	}
	svc := servicediscovery.NewFromConfig(cfg)
	return &ECSServiceDiscoveryClient{
		namespace: namespace,
		svc:       svc,
	}, nil
}

func (e *ECSServiceDiscoveryClient) GetInstanceUrl(ctx context.Context, params GetInstanceUrlInput) (*string, error) {
	if params.ServiceName == "" {
		return nil, errors.New("service name cannot be empty")
	}

	// Discover instances
	input := &servicediscovery.DiscoverInstancesInput{
		NamespaceName: aws.String(e.namespace),
		ServiceName:   aws.String(params.ServiceName),
		MaxResults:    aws.Int32(1),
	}

	result, err := e.svc.DiscoverInstances(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	instance_ip := result.Instances[0].Attributes["AWS_INSTANCE_IPV4"]
	instance_port := result.Instances[0].Attributes["AWS_INSTANCE_PORT"]
	url := fmt.Sprintf("http://%s:%s", instance_ip, instance_port)
	return &url, nil
}
