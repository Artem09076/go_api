package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Artem09076/go_api.git/internal/db/sqlc"
	protos "github.com/Artem09076/go_api.git/protos/gen"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("postgres", "postgresql://postgres:weichahB1eB7zoo@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	srv := grpc.NewServer()

	svc := NewContactServer(queries)

	protos.RegisterContactServiceServer(srv, svc)

	if err := srv.Serve(listener); err != nil {
		log.Fatalln()
	}

}
