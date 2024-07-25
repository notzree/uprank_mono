package main

import (
	"log"
	"net/http"
	"os"

	client "github.com/notzree/uprank_mono/uprank-backend/queue-handler/grpc_client"
	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/queue"
	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/server"
	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/service"
	EnvGetter "github.com/notzree/uprank_mono/uprank-backend/shared/env"
	sd "github.com/notzree/uprank_mono/uprank-backend/shared/service-discovery"
)

func main() {
	//TODO: Based on environment variables, we will switch between different service discovery options
	env := os.Getenv("ENV")
	var eg EnvGetter.EnvGetter
	var sdc sd.ServiceDiscoveryClient

	if env == "dev" {
		log.Default().Println("Running in dev environment")
		eg = EnvGetter.NewAwsEnvGetter("MAIN_BACKEND_SECRETS")
		sdc_pointer, err := sd.NewECSServiceDiscoveryClient("dev.uprank.ca")
		if err != nil {
			log.Fatal("error creating service discovery client:", err)
		}
		sdc = sdc_pointer
	} else {
		log.Default().Println("Running in local environment")
		eg = &EnvGetter.LocalEnvGetter{}
		sdc = sd.NewLocalServiceDiscoveryClient()
	}

	vars := []string{
		"RANKING_QUEUE_URL",
		"MS_API_KEY",
	}
	envVars, err := getEnvVariables(eg, vars)
	if err != nil {
		log.Fatalf("failed to get environment variable: %v", err)
	}
	ranking_queue_url := envVars["RANKING_QUEUE_URL"]
	ms_api_key := envVars["MS_API_KEY"]
	queue := queue.NewSqsQueue(ranking_queue_url)

	inference_server_url := "uprank-inference-backend:50051"
	// _ := context.TODO()
	// inference_server_url, err := sdc.GetInstanceUrl(ctx, sd.GetInstanceUrlInput{
	// 	ServiceName: "inference-backend",
	// })
	// if err != nil {
	// 	log.Fatal("error getting instance url:", err)
	// }
	grpc_inference_client, err := client.NewGRPCInferenceClient(inference_server_url)
	if err != nil {
		log.Fatal("error creating grpc client:", err)
	}
	svc := service.NewUprankVecService(service.NewUprankVecServiceInput{
		ServiceDiscoveryClient: sdc,
		MsApiKey:               ms_api_key,
		InferenceClient:        grpc_inference_client,
		HttpClient:             http.Client{},
	})
	if err != nil {
		log.Fatal("error creating service:", err)
	}

	go func() {
		http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("healthy"))
		})
		log.Fatal(http.ListenAndServe(":80", nil))
	}()

	server := server.NewServer(queue, *svc)
	server.Start()
}

func getEnvVariables(env_getter EnvGetter.EnvGetter, vars []string) (map[string]string, error) {
	envVars := make(map[string]string)
	for _, v := range vars {
		value, err := env_getter.GetEnv(v)
		if err != nil {
			return nil, err
		}
		envVars[v] = value
	}
	return envVars, nil
}
