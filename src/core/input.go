package core

import (
	"cli-book/src/generator"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)
type model struct {
	input textinput.Model
	err   error
}

func newModel() model {
	ti := textinput.New()
	ti.Placeholder = "Name"
	ti.Focus()
	return model{
		input: ti,
		err:   nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			generator.CreateFolder(m.input.Value())
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
func (m model) View() string {
	return fmt.Sprintf(
		"What’s your name app?\n\n%s\n\n%s",
		m.input.View(),
		"(esc to quit)",
	) + "\n"
}

func InitCli() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Has error:%v", err)
		os.Exit(1)
	}
}
