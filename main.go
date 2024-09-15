package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	// "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}


	router := chi.NewRouter() // creating new router object

	// setting the config headers for the router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	// creating a server 
	srv := &http.Server {
		Handler: router,
		Addr: ":" + portString,
	}

	fmt.Printf("server starting on port %v\n", portString)

// HINT srv.ListenAndServe() return and error if the request is not successful so we can use log.Fatal to log the error and using  the error  like  a global error handle
	err := srv.ListenAndServe() // starting the server
	if err != nil {
		log.Fatal(err)
	}


	return
}

