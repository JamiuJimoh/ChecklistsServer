package main

import (
	"log"

	"github.com/JamiuJimoh/checklist/internal/server"
)

func main() {
	server := server.NewServer()

	log.Fatal(server.ListenAndServe())
}
