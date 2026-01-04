# IP-Enrich

[![CI](https://github.com/dalryan/ip-enrich/actions/workflows/build.yml/badge.svg)](https://github.com/dalryan/ip-enrich/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dalryan/ip-enrich)](https://goreportcard.com/report/github.com/dalryan/ip-enrich)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/dalryan/ip-enrich)](https://github.com/dalryan/ip-enrich/blob/main/go.mod)

**A very fast threat intel aggregator.**

IP-Enrich takes a single IP target and concurrently fetches real-time intel from public sources.

## Install

### Binary (Recommended)
Download a pre-compiled binary for your OS from the [Releases Page](https://github.com/dalryan/ip-enrich/releases).

### Go install (Also Recommended)
```shell
go install github.com/dalryan/ip-enrich@v0.1.0
```

### Build from source
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
