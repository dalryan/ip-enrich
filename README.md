# IP-Enrich

IP-Enrich is a Go-based TUI that provides a quick and dirty enrichment of a single IP Address by aggregating data from multiple API endpoints. 
It requires no API keys or local DBs. 

![Demo](assets/demo.gif)

## Usage

To get started, you need to have [Go](https://go.dev/) installed. Once installed, follow these steps:

1. Clone the repository: `git clone https://github.com/dalryan/ip-enrich.git`
2. Navigate to the project directory: `cd ip-enrich`
3. Build the project: `go build`
4. Run the project: `./ip-enrich <ip>`


## Adding New API Endpoints

To add a new API endpoint, you need to modify the `Endpoints` variable in the `model.go` file. Each endpoint is represented as a `APIQueryUnit` which includes the name of the endpoint, the URL, and the model representing the response.

Currently, it assumes the endpoint you are adding requires only a simple GET request. If the endpoint requires additional headers, or a different HTTP method, you will need to modify the `api.go` file to handle these requirements.

## Roadmap

(aka things I will probably never do)

- [ ] Add a summary of the results in the choices view.
- [ ] Add a JSON export of the results.
- [ ] Add support for stdin/stdout piping (e.g. `echo "127.0.0.1" | ip-enrich - | jq ..` )
- [ ] Add more API endpoints for IP enrichment.
- [ ] Add support for host/domain enrichment.
- [ ] Add support for local DB and file lookups.
- [ ] Add some tests? 

## Reference

The project uses the following libraries:

- [Bubbletea](https://github.com/charmbracelet/bubbletea): For the MVU pattern.
- [Lipgloss](https://github.com/charmbracelet/lipgloss): For styling the interface.
- [Bubbles](https://github.com/charmbracelet/bubbles): For additional UI components like the spinner and viewport.
- [Cobra](https://github.com/spf13/cobra): For CLI commands and flags.
- [Chroma](https://github.com/alecthomas/chroma): For syntax highlighting in the JSON view.

The tool was mostly written as an exercise to become familiar with TUIs, Bubbletea, and the MVU pattern. It also takes some inspiration from the excellent [ASN](https://github.com/nitefood/asn)

The tool is intended only for single IP address enrichment and is not intended for bulk enrichment.