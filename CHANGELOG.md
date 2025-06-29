# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2025-06-29

### Added
- Initial Go module structure with proper versioning
- Project infrastructure setup with version management
- Comprehensive development plan (IDEA.md)
- Basic project structure modernization
- Constants files for all packages to eliminate magic numbers
- Apache License 2.0 for open source distribution
- Professional README.md with comprehensive documentation and usage examples
- Project attribution linking to original gnulnx/goperf repository
- **Professional Makefile with comprehensive automation**
  - Build targets for multiple platforms (Linux, macOS, Windows)
  - Load testing automation with predefined test suites
  - Code quality checks (fmt, vet, lint, staticcheck)
  - Dependency management and tool installation
  - Version management and release automation
  - Package building for distribution
  - Integration testing with CLI validation
- **NEW: Clean Architecture Implementation**
  - Core interfaces for HTTPClient, AssetParser, MetricsCollector, OutputFormatter
  - Dependency Injection Container for service management
  - Unified domain models replacing 8+ duplicate structures
  - Multi-source configuration system (flags, env vars, config files)
  - Structured error handling with custom error types
  - Application framework with graceful shutdown
- **NEW: Modern Package Structure**
  - `interfaces/` - Core business abstractions
  - `core/` - Application logic and DI container
  - `cmd/` - Clean command-line entry point
  - Clean separation of concerns and SOLID principles

### Security
- **CRITICAL**: Removed suspicious hardcoded IP address 138.197.97.39 (DigitalOcean server) from CORS configuration
- **CRITICAL**: Replaced non-existent domain qa.teaquinox.com with safe testing API httpbin.org
- Eliminated all magic numbers throughout codebase following zero-tolerance policy
- Added security comments to identify removed threats

### Removed
- Legacy monolithic goperf.go (213 lines) - replaced by clean cmd/main.go (46 lines)  
- Global constants.go - moved to package-specific constant files
- Duplicated mock implementations (core/simple_mocks.go)
- Obsolete binary files and temporary test artifacts
- Unused directories: binaries/, readme_imgs/, .vscode/

### Enhanced
- Comprehensive constants definitions for HTTP status codes, configuration values, and processing limits
- Improved code maintainability with named constants for all numeric literals
- Better documentation with safe example URLs
- Streamlined project structure with 21 Go files (down from 30+)
- Consolidated mock implementations in core/container.go  
- Updated documentation with modern Go development practices
- **NEW: Architecture Improvements**
  - Testable and mockable interfaces for all core services
  - Flexible configuration management with environment variable support
  - Extensible design for new features (WebSocket testing, multiple output formats)
  - Professional error handling with context and wrapping
- Improved Package Building Documentation with comprehensive testing procedures
- Code quality improvements with comprehensive linter fixes achieving zero issues
- Enhanced constants management with descriptive named constants and proper GoDoc comments
- Improved import formatting following Go standards (standard → external → local packages)
- Enhanced code documentation with proper exported constant comments

### Fixed
- **NEW: Architecture Issues**
  - Monolithic main function (130+ lines) replaced with clean App structure
  - Eliminated tight coupling between packages
  - Replaced duplicate structures with unified domain models

### Technical Details
- Initialized Go module for github.com/Gosayram/goperf
- Created .release-version file for version tracking
- Established proper changelog management
- Implemented constants files in all packages (main, httputils, request, perf)
- Replaced all magic numbers with descriptive constants
- Cleaned dependency injection container to use inline mock implementations
- Removed code duplication while maintaining interface compliance
- Improved project maintainability through strategic file organization
- **NEW: Interface-Driven Design**
  - HTTPClient interface abstracts all HTTP operations
  - AssetParser interface supports multiple parsing strategies (regex, DOM, hybrid)
  - MetricsCollector interface provides comprehensive performance measurement
  - OutputFormatter interface supports JSON, text, CSV, HTML formats
- **NEW: Configuration Framework**
  - HTTPConfig, TestConfig, LogConfig, WebConfig, ParserConfig, OutputConfig
  - Environment variable support with GOPERF_ prefix
  - Command-line flag integration with proper defaults
  - Configuration validation and error handling 