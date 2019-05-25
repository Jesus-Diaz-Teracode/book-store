package handlers

import (
	bs "github.com/Jesus-Diaz-Teracode/book-service/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	BC bs.BookClient
}

func (h *Handler) GetBook (c *gin.Context) {
	r, err := h.BC.GetBook(c, &bs.BookRequest{})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error getting the book"})
		return
	}

	c.JSON(http.StatusOK, r)
}
