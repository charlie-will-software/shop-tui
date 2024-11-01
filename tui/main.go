package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuNode struct {
	title    string
	children []*MenuNode
	parent   *MenuNode
}

type Model struct {
	cursor   int
	current  *MenuNode
}

// Define the menu tree structure
func createMenuTree() *MenuNode {
	// Main Menu
	mainMenu := &MenuNode{title: "Main Menu"}

	// Submenus
	viewRecords := &MenuNode{title: "View Records", parent: mainMenu}
	createRecords := &MenuNode{title: "Create Records", parent: mainMenu}
	deleteRecords := &MenuNode{title: "Delete Records", parent: mainMenu}

	// View Records Submenu
	viewAll := &MenuNode{title: "View All", parent: viewRecords}
	searchByID := &MenuNode{title: "Search By ID", parent: viewRecords}
	backFromView := &MenuNode{title: "Back", parent: viewRecords}
	viewRecords.children = []*MenuNode{viewAll, searchByID, backFromView}

	// Create Records Submenu
	addNewRecord := &MenuNode{title: "Add New Record", parent: createRecords}
	importRecords := &MenuNode{title: "Import Records", parent: createRecords}
	backFromCreate := &MenuNode{title: "Back", parent: createRecords}
	createRecords.children = []*MenuNode{addNewRecord, importRecords, backFromCreate}

	// Delete Records Submenu
	deleteByID := &MenuNode{title: "Delete By ID", parent: deleteRecords}
	clearAll := &MenuNode{title: "Clear All Records", parent: deleteRecords}
	backFromDelete := &MenuNode{title: "Back", parent: deleteRecords}
	deleteRecords.children = []*MenuNode{deleteByID, clearAll, backFromDelete}

	// Add submenus to main menu
	mainMenu.children = []*MenuNode{viewRecords, createRecords, deleteRecords}

	return mainMenu
}

func initialModel() Model {
	return Model{
		cursor:  0,
		current: createMenuTree(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.current.children)-1 {
				m.cursor++
			}

		case "enter":
			selected := m.current.children[m.cursor]
			if selected.title == "Back" {
				m.current = m.current.parent
				m.cursor = 0
			} else if len(selected.children) > 0 {
				m.current = selected
				m.cursor = 0
			}

		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	var output string

	// Display the current menu title
	output += fmt.Sprintf("%s:\n", m.current.title)

	// Display each menu item with a cursor
	for i, child := range m.current.children {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // current cursor
		}
		output += fmt.Sprintf("%s %s\n", cursor, child.title)
	}

	output += "\nPress q to quit.\n"
	return output
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

