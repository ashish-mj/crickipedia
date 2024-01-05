package main

import (
	"log"
	"crickipedia/db"
	"crickipedia/handlers"
	env "github.com/joho/godotenv"
	"os"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	
	err := env.Load()
    if err != nil {
    log.Fatalf("Error loading .env file")
    }

	err = db.InitDb(os.Getenv("HOST_NAME"),os.Getenv("BUCKET_NAME"),os.Getenv("USERNAME"),os.Getenv("PASSWORD"),os.Getenv("SCOPE"),os.Getenv("COLLECTION"))
    if err != nil {
		log.Fatalf("Couchbase Connection failed!")
	} else {
		log.Println("Connection Successfull main function")
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/players", handlers.GetAllPlayers).Methods("GET")
	router.HandleFunc("/api/players/{id}", handlers.GetPlayerById).Methods("GET")
	router.HandleFunc("/api/players", handlers.CreatePlayer).Methods("POST")
	router.HandleFunc("/api/players/{id}",handlers.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/api/players/{id}", handlers.UpdatePlayer).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}