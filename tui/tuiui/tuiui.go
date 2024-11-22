package tuiui

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "charlie-will-software/shop-tui/tui/requests"
    "fmt"
    "strconv"
)

type Item struct {
    id int
    name string
    price float64
}

func CreateViewAllItemsPage(pages *tview.Pages) {
    // get items and create table, if non available return
    items := requests.GetItems()
    rows := len(items)
    if rows == 0{
        textView := tview.NewTextView().
        SetText("No Data Found").
        SetDoneFunc(func(key tcell.Key) {
            if (key == tcell.KeyEscape) {
                pages.SwitchToPage("Main Menu") 
            }
        })
        pages.AddAndSwitchToPage("NoDataFoundNotice",textView,false)

    }

    table := createAllItemTable(items,pages)
    pages.AddAndSwitchToPage("ViewAllResponse",table,false)
}

func CreateGetById(pages *tview.Pages){
    getByIdInputField := tview.NewInputField().
    SetLabel("Enter ID Number:").
    SetFieldWidth(19)
    
    textView := tview.NewTextView()

    getByIdInputForm := tview.NewForm().
    AddFormItem(getByIdInputField).
    AddButton("Find Item", func(){
        id_to_get :=getByIdInputField.GetText() 
        item := requests.GetItemById(id_to_get)

        //check if request empty
        if (requests.Item{} == item){
            textView.SetText("No item found for id: " +id_to_get) 
        } else {
            table:= createGetByIdItemTable(item, pages)
            textView.SetText("") 
            pages.AddAndSwitchToPage("ViewAllResponse",table,false)
        }
    }).
    AddButton("Back", func() {
        pages.SwitchToPage("Main Menu")
    }).
    AddFormItem(textView)


    pages.AddPage("Get By Id Form", getByIdInputForm, true, false)
}

func CreateAddItemPage(pages *tview.Pages){
    // Create form to add value to the view

    // Create fields that require reading so they can be referenced within form
    idInputField := tview.NewInputField().
    SetLabel("Id:").
    SetFieldWidth(20).
    SetAcceptanceFunc(tview.InputFieldInteger)

    nameInputField := tview.NewInputField().
    SetLabel("Name:").
    SetFieldWidth(20)

    priceInputField := tview.NewInputField().
    SetLabel("Price:").
    SetFieldWidth(20).
    //TODO: Change acceptance function to only allow 2d.p.
    SetAcceptanceFunc(tview.InputFieldFloat)

    textView := tview.NewTextView()

    // Create Add Form
    addItemForm := tview.NewForm().
    AddFormItem(idInputField).
    AddFormItem(nameInputField).
    AddFormItem(priceInputField).
    AddButton("Add Item",func(){
        id := idInputField.GetText()
        id_int , err_id := strconv.Atoi(id)
        name := nameInputField.GetText()
        price := priceInputField.GetText()
        price_float, err_price := strconv.ParseFloat(price,64)

        if (err_id != nil) {
            textView.SetText("Could not convert id " + id + " to integer")
        } else if (err_price != nil) {
            textView.SetText("Could not convert price " + price + " to float")
        } else{ 
            if (requests.AddItem(id_int, name, price_float)){
                textView.SetText("Item added!")           
            } else {
                textView.SetText("Could not add item.")
            }
        }
    }).
    AddButton("Quit", func(){
        pages.SwitchToPage("Main Menu") 
    }).
    AddFormItem(textView)
    pages.AddPage("Add Item Form", addItemForm, true, false)
}

func CreateDeleteItemPage(pages *tview.Pages){

    idInputField := tview.NewInputField().
    SetLabel("Id:").
    SetFieldWidth(20).
    SetAcceptanceFunc(tview.InputFieldInteger)

    textView := tview.NewTextView()

    deleteItemForm := tview.NewForm().
    AddFormItem(idInputField).
    AddButton("Check For Item",func(){
        id := idInputField.GetText()
        if (requests.Item{} == requests.GetItemById(id)){
            textView.SetText("Item "+ id + " is not available")
        } else {
            textView.SetText("Item " + id + " exists")
        }
    }).
    AddButton("Delete", func() {
        id := idInputField.GetText()
        if (requests.DeleteItem(id)){
            textView.SetText("Item with id " + id + " deleted")
        } else{
            textView.SetText("Could not delete item with id " + id)
        }
    }).
    AddButton("Quit", func(){
        textView.SetText("")
        pages.SwitchToPage("Main Menu")
    }).
    AddFormItem(textView)
    pages.AddPage("Delete Item Form", deleteItemForm, true, false)
}

func createAllItemTable(items []requests.Item, pages *tview.Pages) (*tview.Table) {

    table := createEmptyItemTable(pages)

    // add to items table
    for r := 0; r < len(items); r++ {
        curr_row := items[r]
        addItemRowToTable(table,curr_row,r)
    }

    return table
}

func createGetByIdItemTable(item requests.Item,pages *tview.Pages) (*tview.Table) {
    table:= createEmptyItemTable(pages)

    addItemRowToTable(table, item, 0)

    return table
}

func createEmptyItemTable(pages *tview.Pages) (*tview.Table){

    table := tview.NewTable().SetBorders(true)
    table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
        if (key == tcell.KeyEscape) {
            pages.SwitchToPage("Main Menu") 
        } else if (key == tcell.KeyEnter) {
            table.SetSelectable(true, true)
        }
    }).SetSelectedFunc(func(row int, column int) {
        table.GetCell(row, column).SetTextColor(tcell.ColorRed)
        table.SetSelectable(false, false)
    })


    return table
}

func addItemRowToTable(table *tview.Table, item_to_add requests.Item, row int){
    table.SetCell(row, 0,
    tview.NewTableCell(strconv.Itoa(item_to_add.Id)).
    SetAlign(tview.AlignCenter))
    table.SetCell(row, 1,
    tview.NewTableCell(string(item_to_add.Title)).
    SetAlign(tview.AlignCenter))
    table.SetCell(row, 2,
    tview.NewTableCell(fmt.Sprintf("%.2f",item_to_add.Price)).
    SetAlign(tview.AlignCenter))

}
