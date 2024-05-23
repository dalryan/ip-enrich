package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dalryan/ip-enrich/internal/api"
	"github.com/dalryan/ip-enrich/internal/sources"
	"github.com/dalryan/ip-enrich/internal/utils"
)

var Endpoints = []sources.APIQueryUnit{
	{Name: "Shodan", URL: "https://internetdb.shodan.io/{ip}", Model: &sources.ShodanResponse{}},
	{Name: "GreyNoise", URL: "https://api.greynoise.io/v3/community/{ip}", Model: &sources.GreyNoiseCommunityResponse{}},
	{Name: "IP API", URL: "https://api.ipapi.is/?q={ip}", Model: &sources.IPAPIResponse{}},
	{Name: "IP Whois", URL: "https://ipwho.is/{ip}", Model: &sources.IPInfoResponse{}},
	{Name: "Stop Forum Spam", URL: "https://api.stopforumspam.org/api?json&ip={ip}", Model: &sources.StopForumSpamResponse{}},
}

type Model struct {
	Results         map[string]sources.Result
	Done            int
	Summary         string
	Choice          int
	ViewingResponse bool
	Spinner         spinner.Model
	Viewport        viewport.Model
	Ready           bool
	Ip              string
}

func (m Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, len(Endpoints))
	for i := range Endpoints {
		endpoint := Endpoints[i]
		cmds = append(cmds, api.MakeAPIRequest(endpoint.Name, endpoint.URL, endpoint.Model))
	}
	cmds = append(cmds, m.Spinner.Tick)
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case sources.StatusMsg:
		m.Results[msg.KEY] = sources.Result{
			Data: msg.DATA,
			Url:  msg.URL,
			Code: msg.Code,
			Name: msg.KEY,
			Done: true,
		}
		m.Done++

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.Viewport.YPosition = headerHeight
			m.Ready = true
			m.Viewport.YPosition = headerHeight + 1
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height - verticalMarginHeight
		}

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if !m.ViewingResponse {
				m.ViewingResponse = true
				chosen := Endpoints[m.Choice]
				jsonData, err := utils.PrettyPrintJSON(m.Results[chosen.Name].Data)
				// will need to decide how to handle these errors if they arise
				if err != nil {
					return m, nil
				}
				highlightedJSON, err := utils.HighlightJSON(jsonData)
				if err != nil {
					return m, nil
				}
				str := lipgloss.NewStyle().Width(m.Viewport.Width).Render(highlightedJSON)
				m.Viewport.SetContent(str)
			}
		case tea.KeyUp:
			if !m.ViewingResponse {
				m.Choice = (m.Choice - 1 + len(Endpoints)) % len(Endpoints)
			}
		case tea.KeyDown:
			if !m.ViewingResponse {
				m.Choice = (m.Choice + 1 + len(Endpoints)) % len(Endpoints)
			}
		}

		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "b":
			if m.ViewingResponse {
				m.ViewingResponse = false
			}
			return m, nil
		}
	}
	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
