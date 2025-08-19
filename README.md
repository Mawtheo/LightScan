# LightScan

LightScan is a simple and fast TCP port scanner written in Go. It allows you to quickly scan target hosts for open TCP ports.

## Features

- Fast concurrent scanning
- Customizable port ranges
- Simple command-line interface

## Usage

```bash
go run main.go
```

## Requirements

- Go 1.24.5 or later

## Installation

Clone the repository and build:

```bash
git clone https://github.com/Mawtheo/LightScan.git
cd LightScan
go run main.go
```

## License

MIT License

---

*Currently supports only TCP port scanning. UDP and other features may be added in the future.*