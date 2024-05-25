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
	"github.com/notzree/uprank-backend/main-backend/ent"
)

func main() {
	clerk_secret_key := os.Getenv("CLERK_SECRET_KEY")
	db_connection_string := os.Getenv("DB_CONNECTION_STRING")
	server_port := os.Getenv("SERVER_PORT")
	scraper_queue_url := os.Getenv("SCRAPER_QUEUE_URL")

	//Create db connection
	client, err := ent.Open("postgres", db_connection_string)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// create sqs session
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	sqs_client := sqs.NewFromConfig(sdkConfig)

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

	server := api.NewServer(server_port, router, client, clerk_secret_key, scraper_queue_url, sqs_client)
	fmt.Println("Server listening on port:", server_port)
	log.Fatal(server.Start())
}
