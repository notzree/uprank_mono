package servicediscovery

import (
	"context"
)

type ServiceDiscoveryClient interface {
	GetInstanceUrl(ctx context.Context, params GetInstanceUrlInput) (*string, error)
}

type GetInstanceUrlInput struct {
	ServiceName string
}
