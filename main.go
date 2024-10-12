package main

import (
	"fmt"
	"os"

	"github.com/Carter907/go-4sodoku/consts"
	"github.com/Carter907/go-4sodoku/solve"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	Rows       = consts.Rows
	Cols       = consts.Cols
	Sparseness = consts.Sparseness
)

type model struct {
	board  [Rows][Cols]int // items on the to-do list
	cursor [2]int          // which to-do list item our cursor is pointing at
	solved bool
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	possibleChoices := [Rows]int{}
	for i := range Rows {
		possibleChoices[i] = i + 1
	}
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
			if m.cursor[0] < Rows-1 {
				m.cursor[0]++
			}
		case "left", "h":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}

		case "right", "l":
			if m.cursor[1] < Cols-1 {
				m.cursor[1]++
			}
		case "enter":

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := fmt.Sprintf("Solve the %d by %d Sudoku\n\n", Rows, Cols)

	// Iterate over our choices
	for i, row := range m.board {
		for j, num := range row {
			if num == -1 {
				s += " [ ]"
			} else {
				if i == m.cursor[0] && j == m.cursor[1] {
					s += " [<]"
				} else {
					s += fmt.Sprintf(" [%d]", num)
				}
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
	board := [Rows][Cols]int{}
	solve.Solve(&board, 0, 0)
	solve.Sparcify(&board, Sparseness)

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
