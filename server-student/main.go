package main

import (
	"log"
	"net"

	"github.com/Arturo0911/students-test-microservices/database"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
}
