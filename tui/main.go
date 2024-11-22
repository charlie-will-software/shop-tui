package main

import (
    "log"
    "github.com/rivo/tview"
    "charlie-will-software/shop-tui/tui/tuiui"

)

var pages = tview.NewPages()
var app = tview.NewApplication() //initialise


func main() { 

    tuiui.CreateAddItemPage(pages)
    tuiui.CreateDeleteItemPage(pages)
    tuiui.CreateGetById(pages)

    // create initial list view
    list := tview.NewList().
        AddItem("View All", "View items ordered by ID", 'a',func() {
            tuiui.CreateViewAllItemsPage(pages)
        }).
        AddItem("Get Item By Id", "View singular item", 'b', func() {
            pages.SwitchToPage("Get By ID")
        }).
        AddItem("Add Item", "Add item to database", 'c', func() {
            pages.SwitchToPage("Add Item Form")
        }).
        AddItem("Delete Item", "Delete item from database", 'd', func() {
            pages.SwitchToPage("Delete Item Form")
        }).
        AddItem("Quit", "Press to exit", 'q', func() {
            app.Stop()
        })

    // Add menu items to pages so they can be cycled between
    pages.AddPage("Main Menu", list, true, true)

    // Set the grid as the application root and focus the input field
    app.SetRoot(pages, true).EnableMouse(true)
    
    // Run the application;
    err := app.Run()
    if err != nil {
        log.Fatal(err)
    }

}

