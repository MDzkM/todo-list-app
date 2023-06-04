package main

import (
	"fmt"
	"log"

	"../../internal/router"
)

func main() {
	r := router.Router()
	port := ":8080"
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(r.Listen(port))
}
