package main

import (
	"context"
	"log"
	"time"

	protos "github.com/Artem09076/go_api.git/protos/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient("127.0.0.1:10000", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Panicln(err)
		}
	}()

	client := protos.NewContactServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	createReq := &protos.CreateContactRequest{
		Username: "John Doe",
		Email:    "johndoe@example.com",
	}

	contact, err := client.CreateContact(ctx, createReq)
	if err != nil {
		log.Fatalf("could not create contact: %v", err)
	}

	log.Printf("Contact created: %v", contact)
}
