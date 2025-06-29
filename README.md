# GoPerf - Modern Load Testing Framework

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.24-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-success)](https://github.com/Gosayram/goperf)

A high-performance, concurrent website load testing tool with clean architecture and extensible design.

> **This is a modernized fork** of the original [gnulnx/goperf](https://github.com/Gosayram/goperf) project, redesigned with clean architecture principles, dependency injection, and modern Go practices.

## Project Status

- **Original Project**: [gnulnx/goperf](https://github.com/Gosayram/goperf)
- **This Fork**: [Gosayram/goperf](https://github.com/Gosayram/goperf)

## What's New in This Fork

### ✨ Clean Architecture
- **Dependency Injection Container** - Modular service management
- **Interface-Driven Design** - Easily extensible components
- **SOLID Principles** - Maintainable and testable codebase
- **Zero Magic Numbers** - All constants properly defined

### 🔧 Modern Infrastructure
- **Go Modules** - Modern dependency management
- **Semantic Versioning** - Proper release management
- **Comprehensive Testing** - Unit tests and benchmarks
- **Linter Integration** - Code quality assurance

### 🚀 Enhanced Features
- **Multi-Source Configuration** - CLI flags, environment variables, config files
- **Multiple Output Formats** - JSON, CSV, HTML, plain text
- **Graceful Shutdown** - Proper resource cleanup
- **Structured Logging** - Better debugging and monitoring

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/Gosayram/goperf.git
cd goperf

# Build the application
go build -o goperf ./cmd/main.go
```

### Basic Usage

#### Simple Load Test
```bash
# Test a website with 10 concurrent users for 30 seconds
./goperf -url https://httpbin.org/get -users 10 -sec 30

# Test with custom configuration
./goperf -url https://example.com -users 50 -sec 60 -timeout 10s
```

#### Fetch and Analyze
```bash
# Fetch a page and display detailed information
./goperf -url https://httpbin.org/get -fetch

# Fetch with JSON output including all assets
./goperf -url https://example.com -fetch -json
```

#### Web Interface
```bash
# Start web server mode for browser-based testing
./goperf -web -port 8080
```

## Architecture Overview

This modernized version follows clean architecture principles:

```
├── cmd/                    # Application entry point
├── core/                   # Core application logic
│   ├── app.go             # Main application coordinator  
│   ├── config.go          # Multi-source configuration
│   └── container.go       # Dependency injection
├── interfaces/            # Business logic contracts
│   ├── client.go          # HTTP client interface
│   ├── parser.go          # Asset parser interface
│   ├── metrics.go         # Metrics collection interface
│   └── formatter.go       # Output formatting interface
├── implementations/       # Interface implementations
├── httputils/            # HTTP utilities (legacy)
├── perf/                 # Performance testing (legacy)
└── request/              # Request handling (legacy)
```

## Configuration Options

### Command Line Flags
```bash
-url string         Target URL for testing
-users int          Number of concurrent users (default: 1)
-sec int            Test duration in seconds (default: 10)
-fetch              Fetch mode - analyze single request
-web                Start web server mode
-port int           Web server port (default: 8080)
-timeout duration   Request timeout (default: 30s)
-json               Output in JSON format
-verbose            Enable verbose logging
```

### Environment Variables
```bash
GOPERF_URL=https://example.com
GOPERF_USERS=50
GOPERF_DURATION=60
GOPERF_TIMEOUT=30s
```

### Configuration File
```yaml
# goperf.yaml
target:
  url: "https://example.com"
  timeout: "30s"
load:
  users: 50
  duration: 60
output:
  format: "json"
  verbose: true
```

## Development

### Running Tests
```bash
# Run all tests with coverage
go test ./... -cover

# Run benchmarks
go test ./... -bench=. -benchmem

# Run linter
golangci-lint run
```

### Building for Multiple Platforms
```bash
# Use the provided build script
./build.sh

# Or build manually
GOOS=linux GOARCH=amd64 go build -o goperf-linux-amd64 ./cmd/main.go
GOOS=windows GOARCH=amd64 go build -o goperf-windows-amd64.exe ./cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o goperf-darwin-amd64 ./cmd/main.go
```

## Features

### Core Capabilities
- **High Concurrency** - Leverages Go goroutines for maximum performance
- **Real Browser Simulation** - Fetches CSS, JavaScript, and image assets
- **Session Management** - Maintains cookies across requests
- **Asset Analysis** - Detailed breakdown of page resources
- **Flexible Output** - Multiple output formats for different use cases

### Load Testing Metrics
- Total requests and success rate
- Average, minimum, and maximum latency
- Requests per second (throughput)
- Total bytes transferred
- Detailed timing breakdowns

## Contributing

We welcome contributions! Please see our [development plan](IDEA.md) for current priorities.

### Development Setup
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linter
5. Submit a pull request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Original project by [gnulnx](https://github.com/Gosayram/goperf)
- Clean architecture implementation inspired by modern Go best practices
- Community feedback and contributions

---

**Note**: This is a significant rewrite of the original GoPerf project with focus on maintainability, extensibility, and modern Go development practices.
