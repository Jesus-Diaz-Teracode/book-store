package main

import (
	svc "github.com/Jesus-Diaz-Teracode/book-service/grpc"
	"github.com/Jesus-Diaz-Teracode/book-store/handlers"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "book-service:50051"
)

func setupEngine(bsc svc.BookClient) *gin.Engine {
	r := gin.Default()
	h := &handlers.Handler{BC: bsc}

	r.GET("/ping")
	r.GET("/books", h.GetBook)

	return r
}

func main() {
	// Creates a book service client
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	bsc := svc.NewBookClient(conn)

	engine := setupEngine(bsc)
	err = engine.Run(":8080")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
}
