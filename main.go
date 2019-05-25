package main

import (
	svc "github.com/Jesus-Diaz-Teracode/book-service/grpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func setupEngine(bc svc.BookClient) *gin.Engine {
	r := gin.Default()

	r.GET("/ping")
	r.GET("/books", func(c *gin.Context) {
		r, err := bc.GetBook(c, &svc.BookRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		c.JSON(200, r)
	})

	return r
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	bc := svc.NewBookClient(conn)

	engine := setupEngine(bc)
	err = engine.Run(":8080")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
}
