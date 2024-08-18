package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", index)

	router.Run("0.0.0.0:8080")
}

func index(c *gin.Context) {
	jsonData := []byte(`{"msg": "Hello world!"}`)

	c.Data(http.StatusOK, "application/json", jsonData)
}
