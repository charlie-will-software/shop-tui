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

    // get items, if non available return
    items := requests.GetItems()
    rows := len(items)
    if rows == 0{
        return
    }

    //create table
    table := tview.
    NewTable().
    SetBorders(true).
    SetDoneFunc(func(key tcell.Key) {
        pages.SwitchToPage("Main Menu")
    })

    // add to items table
    for r := 0; r < rows; r++ {
        curr_row := items[r]
        table.SetCell(r, 0,
        tview.NewTableCell(strconv.Itoa(curr_row.Id)).//fmt.Sprintf("%f",curr_row.Id)).
        SetAlign(tview.AlignCenter))
        table.SetCell(r, 1,
        tview.NewTableCell(string(curr_row.Title)).
        SetAlign(tview.AlignCenter))
        table.SetCell(r, 2,
        tview.NewTableCell(fmt.Sprintf("%.2f",curr_row.Price)).
        SetAlign(tview.AlignCenter))
    }


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


    pages.AddAndSwitchToPage("ViewAllResponse",table,false)
}


func CreateAddItemPage(pages *tview.Pages){

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
    SetAcceptanceFunc(tview.InputFieldFloat)


    // Create Form
    addItemForm := tview.NewForm().
    AddFormItem(idInputField).
    AddFormItem(nameInputField).
    AddFormItem(priceInputField).
    AddButton("Save",func(){
        id := idInputField.GetText()
        name := nameInputField.GetText()
        price := priceInputField.GetText()
        fmt.Printf("Form Data:\nID: %s\nName: %s\nPrice: %s\n", id, name, price)
    }).
    AddButton("Quit", func(){
        pages.SwitchToPage("Main Menu") 
    })      
    pages.AddPage("Add Item Form", addItemForm, true, false)
}

func CreateDeleteItemPage(pages *tview.Pages){
    deleteItemForm := tview.NewForm().
    AddInputField("ID to delete", "", 19, nil, nil).
    AddButton("Check For Item",nil).
    AddButton("Delete", nil).
    AddButton("Quit", func(){
        pages.SwitchToPage("Main Menu") 
    })
    pages.AddPage("Delete Item Form", deleteItemForm, true, false)
}

func CreateGetById(pages *tview.Pages){
    getByIdInputField := tview.NewInputField().
    SetLabel("Enter ID Number:").
    SetFieldWidth(19)

    getByIdInputForm := tview.NewForm().
    AddFormItem(getByIdInputField).
    AddButton("Find Item", func(){
        item := request.GetItemById(getByIDInputField.GetText())
        
    }).
    AddButton("Back", func() {
        pages.SwitchToPage("Main Menu")
    })


    pages.AddPage("Get By Id", getByIdInputForm, true, false)
}
