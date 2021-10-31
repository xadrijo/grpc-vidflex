package main

import (
	"log"

	"github.com/xadrijo/grpc-vidflex/internal/db"
	"github.com/xadrijo/grpc-vidflex/internal/product"
)

func Run() error {
	// Initialize and start our gRPC server
	Storage, err := db.New()
	if err != nil {
		return err
	}

	_ = product.New(Storage)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
