# IP-Enrich

**A very fast threat intel aggregator.**

IP-Enrich takes a single IP target and concurrently fetches real-time intel from public sources.

## Install

### Via Go (Recommended)

```shell
go install github.com/dalryan/ip-enrich@latest
```

### From Source
```shell
git clone https://github.com/dalryan/ip-enrich.git
cd ip-enrich
go build -o ip-enrich .
```

## Usage

![Demo](assets/demo.gif)

### Basic scan

Scan an IP using all providers

```shell
ip-enrich 8.8.8.8
```

### List all providers
```shell
ip-enrich list
```

### Advanced Filtering

Scan an IP using only specific providers (comma-separated):

```shell
ip-enrich 1.1.1.1 --providers shodan,greynoise
```

### Automation & Piping

Outputs valid, raw JSON for use with tools like jq:

```shell
ip-enrich 1.1.1.1 --output json | jq '.results[] | select(.status_code == 200)'
```

## Supported providers:
- shodan
- ipapi
- ipwhois
- stopforumspam
- greynoise


## Roadmap

### Features
- [ ] Add support for domain translation
- [ ] Add support for API Keys / Tokens
- [ ] Add support for bulk enrichment
- [ ] Add support for local DB integration
- [ ] Add an optional "summary" 

### Providers
- [ ] BGPView API
- [ ] Team Cymru


## Disclaimer & Responsible Use

- Be responsible when using this tool.
- Always respect the rate-limits and ToS defined by the downstream services. Don't abuse them.
