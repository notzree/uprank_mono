package env

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type EnvGetter interface {
	GetEnv(env string) (string, error)
}

type AwsEnvGetter struct {
	EnvMap map[string]string
}

func NewAwsEnvGetter(NestedSecret string) *AwsEnvGetter {
	all_secrets := os.Getenv(NestedSecret)
	em := make(map[string]string)
	if err := json.Unmarshal([]byte(all_secrets), &em); err != nil {
		log.Fatalf("failed to unmarshal secrets: %v", err)
	}
	return &AwsEnvGetter{
		EnvMap: em,
	}
}

func (a *AwsEnvGetter) GetEnv(env string) (string, error) {
	if val, ok := a.EnvMap[env]; ok {
		return val, nil
	}
	return "", errors.New("env not found")
}

type LocalEnvGetter struct{}

func (l *LocalEnvGetter) GetEnv(env string) (string, error) {
	val := os.Getenv(env)
	if val == "" {
		return "", errors.New("env not found")
	}
	return val, nil
}
