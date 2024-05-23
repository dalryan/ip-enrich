package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"net/http"
	"strings"
)

var (
	baseStyle = lipgloss.NewStyle().Padding(0)

	footerStyle = baseStyle.Copy().
			Foreground(lipgloss.Color("244")).
			Background(lipgloss.Color("238")).
			Bold(true)

	titleStyle = baseStyle.Copy().Bold(true).Foreground(lipgloss.Color("5")).Padding(0, 1)

	subtleStyle = baseStyle.Copy().Foreground(lipgloss.Color("241"))

	keyStyle = baseStyle.Copy().Bold(true).Background(lipgloss.Color("235")).Foreground(lipgloss.Color("254"))

	normalStyle = baseStyle.Copy().Foreground(lipgloss.Color("243"))

	choiceStyle = baseStyle.Copy().Foreground(lipgloss.Color("6"))

	highlightStyle = choiceStyle.Copy().Foreground(lipgloss.Color("205")).Bold(true)
)

func (m Model) headerView() string {
	title := titleStyle.Render("Full Response")
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {

	info := subtleStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m Model) View() string {
	if m.ViewingResponse {
		return lipgloss.NewStyle().PaddingLeft(2).Render(responseView(m))
	} else {
		return lipgloss.NewStyle().PaddingLeft(2).Render(choicesView(m))
	}
}

func responseView(m Model) string {
	chosen := Endpoints[m.Choice]
	info := footerStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	header := titleStyle.Render(fmt.Sprintf("Full Response - %s", chosen.Name))

	footer := lipgloss.JoinHorizontal(lipgloss.Top,
		keyStyle.Render("↑/↓"),
		normalStyle.Render(" scroll "),
		keyStyle.Render("b"),
		normalStyle.Render(" back "),
		keyStyle.Render("q/esc"),
		normalStyle.Render(" quit"),
		normalStyle.Render(info),
	)

	return lipgloss.JoinVertical(lipgloss.Top,
		header,
		m.Viewport.View(),
		footer,
	)
}

func choicesView(m Model) string {
	header := titleStyle.Render(fmt.Sprintf("IP-Enrich - %s", m.Ip) + "\n")
	choices := ""
	for i, endpoint := range Endpoints {
		res := m.Results[endpoint.Name]
		isSelected := m.Choice == i
		choices += checkbox(res.Name, http.StatusText(res.Code), isSelected, m) + "\n"
	}

	instructions := lipgloss.JoinHorizontal(lipgloss.Top,
		keyStyle.Render("↑/↓"),
		normalStyle.Render(" navigate "),
		keyStyle.Render("enter"),
		normalStyle.Render(" select "),
		keyStyle.Render("q/esc"),
		normalStyle.Render(" quit"),
	)

	return lipgloss.JoinVertical(lipgloss.Left, header, choices, instructions)
}

func checkbox(label string, status string, checked bool, m Model) string {
	if status == "" {
		status = m.Spinner.View()
	}
	itemStyle := choiceStyle
	if checked {
		itemStyle = highlightStyle
	}
	checkmark := "[ ]"
	if checked {
		checkmark = "[x]"
	}
	return itemStyle.Render(fmt.Sprintf("%s %s - %s", checkmark, label, status))
}
