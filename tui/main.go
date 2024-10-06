package main

import (
	"fmt"
	"os"
  "log"

	tea "github.com/charmbracelet/bubbletea"

  "charlie-will-software/shop-tui/tui/requests"

)

type model struct {
	choices  []string          // Main menu items
	cursor   int               // Cursor position
	selected map[int]struct{}  // Selected items
	submenu  []string          // Current submenu items
	active   bool              // Whether a submenu is active
}

func initialModel() model {
	return model{
		choices:  []string{"View Records", "Create Records", "Delete Records"},
		selected: make(map[int]struct{}),
		active:   false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			// Exit if we're in the main menu or close the submenu if active
			if m.active {
				m.active = false
				m.cursor = 0 // Reset cursor when returning to main menu
			} else {
				return m, tea.Quit
			}

		case "up", "k":
			// Navigate up in the current menu
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			// Navigate down in the current menu
			if m.cursor < len(m.getCurrentChoices())-1 {
				m.cursor++
			}

		case "enter", " ":
			if m.active {
				// Handle submenu selections
				if m.cursor == len(m.submenu)-1 { // Check if "Back" is selected (last item in submenu)
					m.active = false
					m.cursor = 0 // Reset cursor when returning to the main menu
				} else {
					fmt.Println("Submenu option selected:", m.submenu[m.cursor])
          switch m.submenu[m.cursor] {
          case "View All":
            items, err := requests.GetItems()
            if err != nil {
		          log.Fatalf("Error getting items: %v", err)
	          }
          	fmt.Println("Items:", items)
          }
				}
			} else {
				// Activate submenu based on the main menu selection
				switch m.cursor {
				case 0: // View Records submenu
					m.submenu = []string{"View All", "Search by ID", "Back"}
					m.active = true
				case 1: // Create Records submenu
					m.submenu = []string{"Add New Record", "Import Records", "Back"}
					m.active = true
				case 2: // Delete Records submenu
					m.submenu = []string{"Delete by ID", "Clear All Records", "Back"}
					m.active = true
				}
				m.cursor = 0 // Reset cursor for the submenu
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	if m.active {
		// Render the submenu
		s += "Submenu:\n\n"
		for i, choice := range m.submenu {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	} else {
		// Render the main menu
		s += "Main Menu:\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	}

	s += "\nPress q to quit. Press Enter to select.\n"
	return s
}

// getCurrentChoices returns the current menu items based on the state
func (m model) getCurrentChoices() []string {
	if m.active {
		return m.submenu
	}
	return m.choices
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

