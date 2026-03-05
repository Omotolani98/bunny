package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type field struct {
	label  string
	key    string
	value  string
	masked bool
}

type model struct {
	provider string
	fields   []field
	cursor   int
	done     bool
}

func getProviderFields(provider string) []field {
	switch provider {
	case "hetzner":
		return []field{
			{label: "API Token", key: "token", masked: true},
		}
	case "aws":
		return []field{
			{label: "Access Key ID", key: "access_key_id", masked: false},
			{label: "Secret Access Key", key: "secret_access_key", masked: true},
			{label: "Default Region", key: "region", masked: false},
		}
	default:
		return nil
	}
}

func initialModel(provider string) model {
	return model{
		provider: provider,
		fields:   getProviderFields(provider),
		cursor:   0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.cursor < len(m.fields)-1 {
				m.cursor++
			} else {
				m.done = true
				return m, tea.Quit
			}
		case tea.KeyBackspace:
			f := &m.fields[m.cursor]
			if len(f.value) > 0 {
				f.value = f.value[:len(f.value)-1]
			}
		default:
			if msg.Type == tea.KeyRunes {
				m.fields[m.cursor].value += string(msg.Runes)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("212")).
		Render(fmt.Sprintf("Authenticating with %s", m.provider))

	var b strings.Builder
	b.WriteString(title + "\n\n")

	for i, f := range m.fields {
		display := f.value
		if f.masked && len(display) > 0 {
			display = strings.Repeat("•", len(display))
		}

		if i == m.cursor {
			b.WriteString(fmt.Sprintf("  → %s: %s█\n", f.label, display))
		} else if f.value != "" {
			if f.masked {
				b.WriteString(fmt.Sprintf("  ✓ %s: %s\n", f.label, strings.Repeat("•", len(f.value))))
			} else {
				b.WriteString(fmt.Sprintf("  ✓ %s: %s\n", f.label, f.value))
			}
		} else {
			b.WriteString(fmt.Sprintf("    %s:\n", f.label))
		}
	}

	b.WriteString("\n  Press Enter to continue • Esc to cancel\n")
	return b.String()
}

// Public function to run the prompt and return credentials
func PromptCloudAuth(provider string) (map[string]string, error) {
	m := initialModel(provider)
	final, err := tea.NewProgram(m).Run()
	if err != nil {
		return nil, err
	}

	result := final.(model)
	if !result.done {
		return nil, fmt.Errorf("authentication cancelled")
	}

	credentials := make(map[string]string)
	for _, f := range result.fields {
		if f.value == "" {
			return nil, fmt.Errorf("%s is required", f.label)
		}
		credentials[f.key] = f.value
	}

	return credentials, nil
}
