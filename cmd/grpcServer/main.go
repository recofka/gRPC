package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/recofka/gRPC/internal/database"
	"github.com/recofka/gRPC/internal/pb"
	"github.com/recofka/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcSever := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcSever, categoryService)
	// evans
	reflection.Register(grpcSever)

	// open a tcp connection
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}
	if err := grpcSever.Serve(lis); err != nil {
		panic(err)
	}

}
