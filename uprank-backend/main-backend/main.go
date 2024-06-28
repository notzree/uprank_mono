package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/notzree/uprank-backend/main-backend/api"
	"github.com/notzree/uprank-backend/main-backend/authenticator"
	"github.com/notzree/uprank-backend/main-backend/ent"
	svc "github.com/notzree/uprank-backend/main-backend/service"
)

func main() {
	clerk_secret_key := os.Getenv("CLERK_SECRET_KEY")
	db_connection_string := os.Getenv("DB_CONNECTION_STRING")
	server_port := os.Getenv("SERVER_PORT")
	ranking_queue_url := os.Getenv("RANKING_QUEUE_URL")
	ms_api_key := os.Getenv("MS_API_KEY")

	//Create db connection
	ent_client, err := ent.Open("postgres", db_connection_string)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer ent_client.Close()

	// create sqs session
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	sqs_client := sqs.NewFromConfig(sdkConfig)

	servicer := svc.NewV1Servicer(ent_client, sqs_client, ranking_queue_url)

	authenticator := authenticator.NewClerkAuthenticator(clerk_secret_key, ms_api_key)

	//Create router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*", "chrome-extension://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	server := api.NewServer(server_port, router, authenticator, servicer)
	fmt.Println("Server listening on port:", server_port)
	log.Fatal(server.Start())
}
