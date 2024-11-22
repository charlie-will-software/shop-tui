package requests

import (
    "net/http"
    "time"
    "log"
    "encoding/json"
    "io"
)

var (
    BaseURL = "http://localhost:8080"
    httpClient = &http.Client{
        Timeout: 10 * time.Second, // Change to change timeout for requests
    }
)


type Item struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Price float64 `json:"price"`
}

func GetItems() ([]Item) {
    resp, err := httpClient.Get(BaseURL + "/items")
    if err != nil {
        log.Fatal(err)
    }
    
    var items []Item

    defer resp.Body.Close() // wait till function ends
    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        err_json := json.Unmarshal(bodyBytes, &items)
        if err_json != nil {
            log.Fatal(err)
        }
    }
    return items
}


func GetItemById(id string) (Item){
    resp, err := httpClient.Get(BaseURL + "/items/" + id)
    if err != nil {
        log.Fatal(err)
    } 

    defer resp.Body.Close()

    var item Item

    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        err_json := json.Unmarshal(bodyBytes, &item)
        if err_json != nil {
            log.Fatal(err)
        }

        //TODO: Check Item is only length 1
    }
    return item
    
}


