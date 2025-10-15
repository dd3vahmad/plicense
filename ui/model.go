package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type License struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	SpdxID      string `json:"spdx_id"`
	path        string
	Description string   `json:"description"`
	Body        string   `json:"body"`
	Permissions []string `json:"permissions"`
	Conditions  []string `json:"conditions"`
	Limitations []string `json:"limitations"`
}

func (l License) Title() string       { return l.Name }
func (l License) FilterValue() string { return l.Name }

type model struct {
	list     list.Model
	viewport viewport.Model
	licenses []License
}

func NewModel(dir string, lcs []License) (model, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return model{}, err
	}

	var licenses []License
	for _, f := range files {
		path := filepath.Join(dir, f.Name())
		data, _ := os.ReadFile(path)
		name := f.Name()
		body := string(data)
		licenses = append(licenses, License{Name: name, path: path, Body: body})
	}

	items := make([]list.Item, len(licenses))
	for i, l := range licenses {
		items[i] = l
	}

	licenseList := list.New(items, list.NewDefaultDelegate(), 25, 15)
	licenseList.Title = "Select a License"

	vp := viewport.New(60, 20)
	vp.SetContent(licenses[0].Body)

	return model{list: licenseList, viewport: vp, licenses: licenses}, nil
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			selected := m.list.SelectedItem().(License)
			err := os.WriteFile("LICENSE", []byte(selected.Body), 0o644)
			if err != nil {
				fmt.Println("Failed to write LICENSE:", err)
			} else {
				fmt.Printf("\n Added '%s' license to ./LICENSE\n", selected.Name)
			}
			return m, tea.Quit

		case "u", "ctrl+j":
			m.viewport.ScrollUp(5)
		case "d", "ctrl+k":
			m.viewport.ScrollDown(5)
		case "l", "ctrl+l":
			m.viewport.ScrollLeft(5)
		case "r", "ctrl+h":
			m.viewport.ScrollRight(5)
		}
	}

	m.list, cmd = m.list.Update(msg)
	if sel, ok := m.list.SelectedItem().(License); ok {
		m.viewport.SetContent(sel.Body)
	}

	return m, cmd
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		m.list.View(),
		lipgloss.NewStyle().Margin(0, 2).Border(lipgloss.RoundedBorder()).Padding(2).BorderForeground(lipgloss.Color("205")).Render(m.viewport.View()),
	)
}
