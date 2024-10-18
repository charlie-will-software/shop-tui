package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

  "charlie-will-software/shop-tui/tui/requests/model"
)

var (
  BaseURL   = "http://localhost:8080"
	httpClient = &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for requests
	}
)

// Helper function to check response status code
func checkResponse(resp *http.Response, expectedStatusCode int) error {
	if resp.StatusCode != expectedStatusCode {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// Helper function to decode JSON response
func decodeJSON(body io.ReadCloser, target interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}
	return nil
}

// GetIndex fetches the index page from the base URL
func GetIndex() (string, error) {
	resp, err := httpClient.Get(BaseURL + "/")
	if err != nil {
		return "", fmt.Errorf("failed to fetch index: %w", err)
	}
	defer resp.Body.Close()

	if err := checkResponse(resp, http.StatusOK); err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

// GetItems fetches all items from the "/items" endpoint
func GetItems() ([]model.Item, error) {
	resp, err := httpClient.Get(BaseURL + "/items")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}
	if err := checkResponse(resp, http.StatusOK); err != nil {
		return nil, err
	}

	var items []model.Item
	if err := decodeJSON(resp.Body, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// GetItemByID fetches a specific item by ID from the "/items/:id" endpoint
func GetItemByID(id int) (*model.Item, error) {
	resp, err := httpClient.Get(BaseURL + "/items/" + strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item with ID %d: %w", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("item not found")
	}
	if err := checkResponse(resp, http.StatusOK); err != nil {
		return nil, err
	}

	var item model.Item
	if err := decodeJSON(resp.Body, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

// AddItem sends a POST request to the "/items" endpoint to add a new item
func AddItem(newItem model.Item) (*model.Item, error) {
	itemData, err := json.Marshal(newItem)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal item: %w", err)
	}

	resp, err := httpClient.Post(BaseURL+"/items", "application/json", bytes.NewBuffer(itemData))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to add item: %w", err)
	}
	defer resp.Body.Close()

	if err := checkResponse(resp, http.StatusCreated); err != nil {
		return nil, err
	}

	var createdItem model.Item
	if err := decodeJSON(resp.Body, &createdItem); err != nil {
		return nil, err
	}

	return &createdItem, nil
}

