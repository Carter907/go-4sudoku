package main

import (
	"fmt"
	"os"

	"github.com/Carter907/go-4sodoku/solve"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	BoardRows  = 4
	BoardCols  = 4
	Sparseness = 1
)

type model struct {
	board  [BoardRows][BoardCols]int // items on the to-do list
	cursor [2]int                    // which to-do list item our cursor is pointing at
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor[0] > 0 {
				m.cursor[0]--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor[0] < BoardRows-1 {
				m.cursor[0]++
			}
		case "left", "h":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}

		case "right", "l":
			if m.cursor[1] < BoardCols-1 {
				m.cursor[1]++
			}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Solve the 4 by 4 Sudoku\n\n"

	// Iterate over our choices
	for _, row := range m.board {
		for _, num := range row {
			if num == -1 {
				s += " [ ]"
			} else {
				s += fmt.Sprintf(" [%d]", num)
			}
		}
		// Render the row
		s += "\n"
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func initialModel() model {
	board := [4][4]int{}
	solve.Solve(board, 0, 0)

	return model{
		board: board,
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
