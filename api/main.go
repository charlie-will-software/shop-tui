package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"charlie-will-software/shop-tui/api/model"
)

var items = []model.Item{
	{Id: 1, Title: "Bread", Price: 1.09},
	{Id: 2, Title: "Milk", Price: 1.5},
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/items", getItems)

	router.Run("0.0.0.0:8080")
}

func index(c *gin.Context) {
	jsonData := []byte(`{"msg": "Hello world!"}`)

	c.Data(http.StatusOK, "application/json", jsonData)
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}
