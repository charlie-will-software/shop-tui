package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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
	router.GET("/items/:id", getItemById)
	router.POST("/items", addItem)

	router.Run("0.0.0.0:8080")
}

func index(c *gin.Context) {
	jsonData := []byte(`{"msg": "Hello world!"}`)

	c.Data(http.StatusOK, "application/json", jsonData)
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func getItemById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid item ID"})
		return
	}

	// Find an Item with the correct Id.
	for _, a := range items {
		if a.Id == idInt {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func addItem(c *gin.Context) {
	var newItem model.Item

	if err := c.BindJSON(&newItem); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if newItem.Title == "" || newItem.Price <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid item data"})
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}
