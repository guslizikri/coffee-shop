package main

import (
	"coffee-shop/internal/routers"
	"coffee-shop/pkg"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := pkg.Posql()
	if err != nil {
		log.Fatal(err)
	}

	router := routers.New(db)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
