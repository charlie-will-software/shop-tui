package requests

import (
    "net/http"
    "bytes"
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

func AddItem(id int, title string, price float64) (is_added bool){
    //var item = Item{id,title,price}


    var values = Item{id,title,price}
    json_data, err := json.Marshal(values)

    if err != nil {
        return false //TODO: Add more detailed error messaging
    }

    resp, err := httpClient.Post(BaseURL + "/items", "applicaiton/json", bytes.NewBuffer(json_data))
    defer resp.Body.Close()

    return err == nil

     //TODO add more detailed error messaging
}

func DeleteItem(id string) (is_deleted bool){

    // create request
    req, err := http.NewRequest("DELETE", BaseURL + "/items/" + id,nil)
    if err != nil{
        return false //TODO: add more detailed error messaging
    }

    resp, err:= httpClient.Do(req)
    defer resp.Body.Close()
    if err != nil {
        return false
    }
    return true

}
//TODO: Update Item
