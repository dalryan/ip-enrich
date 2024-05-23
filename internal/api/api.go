package api

import (
	"encoding/json"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dalryan/ip-enrich/internal/sources"
	"io"
	"net/http"
	"time"
)

// MakeAPIRequest sends a GET request to the URL and attempts to decode the JSON response into the supplied model.
// This function is designed to be used with the BubbleTea framework, returning a command (tea.Cmd).
//
// If successful, it returns a statusMsg that includes the URL, the HTTP status code, and the decoded model.
// If unsuccessful, it returns an errMsg containing the URL and the error.
//
// Parameters:
//   - url: The URL to which the HTTP request is sent.
//   - model: A pointer to a struct where the JSON response will be decoded.
//
// Returns:
//   - tea.Cmd: An async command executed by the BubbleTea runtime.
func MakeAPIRequest(key string, url string, model interface{}) tea.Cmd {
	return func() tea.Msg {
		c := &http.Client{Timeout: 10 * time.Second}
		res, err := c.Get(url)
		if err != nil {
			return sources.ErrMsg{URL: url, Err: err}
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(res.Body)
		if err := json.NewDecoder(res.Body).Decode(model); err != nil {
			return sources.ErrMsg{URL: url, Err: err}
		}
		return sources.StatusMsg{URL: url, Code: res.StatusCode, DATA: model, KEY: key}
	}
}
