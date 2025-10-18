package ui

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dd3vahmad/plicense/internals/entity"
	"github.com/dd3vahmad/plicense/internals/fetch"
	"github.com/fatih/color"
)

type model struct {
	list       list.Model
	viewport   viewport.Model
	licenses   []entity.License
	showList   bool
	licenseKey string
}

func NewLicensesListModel(licenses []entity.License) (model, error) {
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

	return model{list: licenseList, viewport: vp, licenses: licenses, showList: true}, nil
}

func NewLicenseModel() (model, error) {
	file, err := os.ReadFile("LICENSE")
	if err != nil {
		fmt.Printf("%s", color.HiYellowString("Seems you're not in the root of your project\n"))
		fmt.Printf("%s", color.HiRedString("Cannot LICENSE file here. Does it exist??\n"))
		return model{}, fmt.Errorf("could not read LICENSE file: %w", err)
	}

	vp := viewport.New(72, 24)
	vp.SetContent(string(file))

	return model{
		viewport: vp,
		showList: false,
	}, nil
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
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

		if m.showList {
			switch msg.String() {
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
				path, _ := fetch.LicensePath(selected.Key)

				newLicense, _ := os.Create(path)
				defer newLicense.Close()

				json.NewEncoder(newLicense).Encode(selected)
				return m, tea.Quit
			}
		}
	}

	if m.showList {
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
	}
	return m, cmd
}

func (m model) View() string {
	if m.showList {
		return lipgloss.JoinHorizontal(lipgloss.Top,
			m.list.View(),
			lipgloss.NewStyle().
				Margin(0, 2).
				Border(lipgloss.RoundedBorder()).
				Padding(2).
				BorderForeground(lipgloss.Color("205")).
				Render(m.viewport.View()),
		)
	}

	return lipgloss.NewStyle().
		Margin(1, 2).
		Border(lipgloss.RoundedBorder()).
		Padding(2).
		BorderForeground(lipgloss.Color("205")).
		Render(m.viewport.View())
}
