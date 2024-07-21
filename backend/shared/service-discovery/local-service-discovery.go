package servicediscovery

import (
	"context"
	"errors"
	"os"
)

type LocalServiceDiscoveryClient struct {
}

func NewLocalServiceDiscoveryClient() *LocalServiceDiscoveryClient {
	return &LocalServiceDiscoveryClient{}
}
func (l *LocalServiceDiscoveryClient) GetInstanceUrl(ctx context.Context, params GetInstanceUrlInput) (*string, error) {
	if params.ServiceName == "" {
		return nil, errors.New("service name cannot be empty")
	}
	if params.ServiceName == "main-backend" {
		url := os.Getenv("MAIN_BACKEND_URL")
		return &url, nil
	}
	if params.ServiceName == "inference-backend" {
		url := os.Getenv("INFERENCE_SERVER_URL")
		return &url, nil
	}
	return nil, errors.New("service not found")
}
