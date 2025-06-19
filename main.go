package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Structure for resource
type Resource struct {
	Name        string
	URL         string
	Description string
}

// List of Go resources
var resources = []Resource{
	{
		Name:        "Go.dev",
		URL:         "https://go.dev",
		Description: "Official Go website - documentation, tutorials, packages",
	},
	{
		Name:        "Go by Example",
		URL:         "https://gobyexample.com",
		Description: "Practical Go code examples with explanations",
	},
	{
		Name:        "Effective Go",
		URL:         "https://go.dev/doc/effective_go",
		Description: "Guide to writing effective Go code",
	},
	{
		Name:        "Go Tour",
		URL:         "https://tour.golang.org",
		Description: "Interactive tour of Go language basics",
	},
	{
		Name:        "Awesome Go",
		URL:         "https://awesome-go.com",
		Description: "Curated list of Go frameworks, libraries and resources",
	},
	{
		Name:        "Codewars",
		URL:         "https://www.codewars.com/dashboard",
		Description: "Programming practice platform",
	},
	{
		Name:        "Exercism",
		URL:         "https://exercism.org/dashboard",
		Description: "Coding exercises and practice",
	},
}

// Application model
type model struct {
	cursor   int
	selected map[int]struct{}
	choice   string
}

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("#7D56F4")).
				Bold(true)

	descriptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#626262")).
				PaddingLeft(6)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Padding(1, 0)
)

// Model initialization
func initialModel() model {
	return model{
		selected: make(map[int]struct{}),
	}
}

// Initialization
func (m model) Init() tea.Cmd {
	return nil
}

// Message handling
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(resources)-1 {
				m.cursor++
			}

		case "enter", " ":
			// Open selected resource in browser
			err := openURL(resources[m.cursor].URL)
			if err != nil {
				m.choice = fmt.Sprintf("Error opening: %v", err)
			} else {
				m.choice = fmt.Sprintf("Opened: %s", resources[m.cursor].Name)
			}
		}

	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseWheelUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.MouseWheelDown:
			if m.cursor < len(resources)-1 {
				m.cursor++
			}
		case tea.MouseLeft:
			// Determine which item was clicked
			if msg.Y >= 3 && msg.Y < 3+len(resources)*3 {
				itemIndex := (msg.Y - 3) / 3
				if itemIndex >= 0 && itemIndex < len(resources) {
					m.cursor = itemIndex
					// Open resource on click
					err := openURL(resources[m.cursor].URL)
					if err != nil {
						m.choice = fmt.Sprintf("Error opening: %v", err)
					} else {
						m.choice = fmt.Sprintf("Opened: %s", resources[m.cursor].Name)
					}
				}
			}
		}
	}

	return m, nil
}

// Display
func (m model) View() string {
	s := titleStyle.Render("🐹 Go Resources - Select a resource to learn Go")
	s += "\n\n"

	for i, resource := range resources {
		cursor := " "
		if m.cursor == i {
			cursor = "►"
			s += selectedItemStyle.Render(fmt.Sprintf("%s %s", cursor, resource.Name))
		} else {
			s += itemStyle.Render(fmt.Sprintf("%s %s", cursor, resource.Name))
		}
		s += "\n"
		s += descriptionStyle.Render(resource.Description)
		s += "\n\n"
	}

	if m.choice != "" {
		s += fmt.Sprintf("\n%s\n", m.choice)
	}

	return s
}

// Function to open URL in browser
func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		// For Windows use start
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		// For macOS use open
		cmd = "open"
		args = []string{url}
	default:
		// For Linux and other Unix systems
		// Try different commands
		browsers := []string{"google-chrome", "chromium-browser", "firefox", "xdg-open"}
		for _, browser := range browsers {
			if _, err := exec.LookPath(browser); err == nil {
				cmd = browser
				args = []string{url}
				break
			}
		}
		if cmd == "" {
			return fmt.Errorf("no suitable browser found")
		}
	}

	return exec.Command(cmd, args...).Start()
}

func main() {
	// Enable mouse support
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),       // Use alternative screen
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	if _, err := p.Run(); err != nil {
		log.Printf("Application startup error: %v", err)
		os.Exit(1)
	}
}
