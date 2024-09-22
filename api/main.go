package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"

	"charlie-will-software/shop-tui/api/model"
)

// Seed the items array with data
var (
	items = []model.Item{
		{Id: 1, Title: "Bread", Price: 1.09},
		{Id: 2, Title: "Milk", Price: 1.5},
	}
	mu sync.Mutex
)

func main() {
	router := gin.Default()

	router.GET("/", index)
	router.GET("/items", getItems)
	router.GET("/items/:id", getItemById)
	router.POST("/items", addItem)
	router.DELETE("/items/:id", deleteItem)

	address := getEnv("SERVER_ADDRESS", "0.0.0.0:8080")
	router.Run(address)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getItemIndexById(id int) (int, error) {
	mu.Lock()
	defer mu.Unlock()
	// Find an Item with the correct Id.
	for i, a := range items {
		if a.Id == id {
			return i, nil
		}
	}

	return 0, errors.New("item with id not found")
}

func index(c *gin.Context) {
	jsonData := []byte(`{"msg": "Hello world!"}`)

	c.Data(http.StatusOK, "application/json", jsonData)
}

func getItems(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.IndentedJSON(http.StatusOK, items)
}

func getItemById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid item ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
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

	mu.Lock()
	items = append(items, newItem)
	mu.Unlock()
	c.IndentedJSON(http.StatusCreated, newItem)
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid item ID"})
		return
	}

	i, err := getItemIndexById(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}

	mu.Lock()
	deletedItem := items[i]
	items = append(items[:i], items[i+1:]...)
	mu.Unlock()

	c.IndentedJSON(http.StatusOK, deletedItem)
}
