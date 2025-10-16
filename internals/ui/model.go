package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dd3vahmad/plicense/internals/entity"
	"github.com/dd3vahmad/plicense/internals/fetch"
)

type model struct {
	list     list.Model
	viewport viewport.Model
	licenses []entity.License
}

func NewModel(licenses []entity.License) (model, error) {
	items := make([]list.Item, len(licenses))
	for i, l := range licenses {
		items[i] = l
	}

	listDelegate := list.NewDefaultDelegate()
	listDelegate.SetHeight(1)
	listDelegate.SetSpacing(1)
	licenseList := list.New(items, listDelegate, 25, 15)
	licenseList.Title = "Select a License"

	vp := viewport.New(72, 24)
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
			selected := m.list.SelectedItem().(entity.License)
			if selected.Body == "" {
				fetched, err := fetch.LicenseDetails(selected.Key)
				if err != nil {
					return m, cmd
				}
				selected.Body = fetched.Body
			}
			err := os.WriteFile("LICENSE", []byte(selected.Body), 0o644)

			if err != nil {
				fmt.Println("Failed to write LICENSE:", err)
			} else {
				fmt.Printf("\n Added '%s' license to ./LICENSE\n", selected.Name)
			}

			// Cache selected license in a JSON file for later use.
			path := filepath.Join("licenses", fmt.Sprintf("%s.json", selected.Key))

			newLicense, _ := os.Create(path)
			defer newLicense.Close()

			json.NewEncoder(newLicense).Encode(selected)
			return m, tea.Quit

		case "t":
			m.viewport.ScrollUp(5)
		case "b":
			m.viewport.ScrollDown(5)
		case "s":
			m.viewport.ScrollLeft(5)
		case "e":
			m.viewport.ScrollRight(5)
		}
	}

	m.list, cmd = m.list.Update(msg)
	if sel, ok := m.list.SelectedItem().(entity.License); ok {
		if sel.Body == "" {
			fetched, err := fetch.LicenseDetails(sel.Key)
			if err != nil {
				m.viewport.SetContent("Error fetching license details")
			} else {
				sel.Body = fetched.Body
				m.viewport.SetContent(sel.Body)
			}
		} else {
			m.viewport.SetContent(sel.Body)
		}
	}

	return m, cmd
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		m.list.View(),
		lipgloss.NewStyle().Margin(0, 2).Border(lipgloss.RoundedBorder()).Padding(2).BorderForeground(lipgloss.Color("205")).Render(m.viewport.View()),
	)
}
