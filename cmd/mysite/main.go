package main

import (
	"fmt"
	"log"
	"webstory/internal/app"
)

const Host = "http://localhost:5000"

func main() {
	server := app.NewServer()

	fmt.Println("Server running on HOST: ", Host)
	log.Fatal(server.Start(":5000"))
}
