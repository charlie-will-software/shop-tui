package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

  "charlie-will-software/shop-tui/tui/requests/model"
  
)

// BaseURL is the URL where the API server is running.
var BaseURL = "http://localhost:8080"

// GetIndex sends a GET request to the root endpoint ("/").
func GetIndex() (string, error) {
	resp, err := http.Get(BaseURL + "/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetItems sends a GET request to the "/items" endpoint and returns all items.
func GetItems() ([]model.Item, error) {
	resp, err := http.Get(BaseURL + "/items")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var items []model.Item
	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// GetItemByID sends a GET request to the "/items/:id" endpoint to fetch a specific item by ID.
func GetItemByID(id int) (*model.Item, error) {
	resp, err := http.Get(BaseURL + "/items/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("item not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var item model.Item
	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// AddItem sends a POST request to the "/items" endpoint to add a new item.
func AddItem(newItem model.Item) (*model.Item, error) {
	itemData, err := json.Marshal(newItem)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(BaseURL+"/items", "application/json", bytes.NewBuffer(itemData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var createdItem model.Item
	err = json.NewDecoder(resp.Body).Decode(&createdItem)
	if err != nil {
		return nil, err
	}

	return &createdItem, nil
}

