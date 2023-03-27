package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get env
	PORT := getEnv("PORT", "3030")
	r, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("app run on port " + PORT)
	err = http.ListenAndServe(":"+PORT, r)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
