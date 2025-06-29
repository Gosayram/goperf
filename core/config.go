package core

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config represents the complete application configuration
// This replaces scattered command-line flags throughout the codebase
type Config struct {
	HTTP   HTTPConfig   `json:"http"`
	Test   TestConfig   `json:"test"`
	Log    LogConfig    `json:"log"`
	Web    WebConfig    `json:"web"`
	Parser ParserConfig `json:"parser"`
	Output OutputConfig `json:"output"`
}

// HTTPConfig contains HTTP client configuration
type HTTPConfig struct {
	Timeout        time.Duration `json:"timeout"`
	MaxConnections int           `json:"max_connections"`
	RetryAttempts  int           `json:"retry_attempts"`
	UserAgent      string        `json:"user_agent"`
}

// TestConfig contains load testing configuration
type TestConfig struct {
	DefaultUsers    int           `json:"default_users"`
	DefaultDuration time.Duration `json:"default_duration"`
	DefaultURL      string        `json:"default_url"`
	OutputFile      string        `json:"output_file"`
	Iterations      int           `json:"iterations"`
	OutputInterval  int           `json:"output_interval"`
}

// LogConfig contains logging configuration
type LogConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
	Output string `json:"output"`
}

// WebConfig contains web server configuration
type WebConfig struct {
	Port    int      `json:"port"`
	CORS    []string `json:"cors"`
	Enabled bool     `json:"enabled"`
	APIPath string   `json:"api_path"`
}

// ParserConfig contains asset parsing configuration
type ParserConfig struct {
	Method     string `json:"method"` // "regex", "dom", "hybrid"
	Concurrent bool   `json:"concurrent"`
	RegexLimit int    `json:"regex_limit"`
}

// OutputConfig contains output formatting configuration
type OutputConfig struct {
	Format      string `json:"format"` // "json", "text", "csv", "html"
	Colors      bool   `json:"colors"`
	Indentation string `json:"indentation"`
}

// LoadConfig loads configuration from multiple sources in order:
// 1. Default values
// 2. Configuration file (if exists)
// 3. Environment variables
// 4. Command line flags
func LoadConfig() (*Config, error) {
	cfg := DefaultConfig()

	// Load from config file (if exists)
	if err := cfg.LoadFromFile(); err != nil {
		// Config file is optional, so just log the error
		fmt.Printf("Warning: Could not load config file: %v\n", err)
	}

	// Load from environment variables
	if err := cfg.LoadFromEnv(); err != nil {
		return nil, fmt.Errorf("failed to load from environment: %w", err)
	}

	// Load from command line flags
	if err := cfg.LoadFromFlags(); err != nil {
		return nil, fmt.Errorf("failed to load from flags: %w", err)
	}

	// Validate the final configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

// DefaultConfig returns configuration with default values
func DefaultConfig() *Config {
	return &Config{
		HTTP: HTTPConfig{
			Timeout:        DefaultHTTPTimeout,
			MaxConnections: DefaultMaxConnections,
			RetryAttempts:  DefaultRetryAttempts,
			UserAgent:      DefaultUserAgent,
		},
		Test: TestConfig{
			DefaultUsers:    DefaultUsers,
			DefaultDuration: DefaultTestDuration,
			DefaultURL:      DefaultURL,
			OutputFile:      DefaultOutputFile,
			Iterations:      DefaultIterations,
			OutputInterval:  DefaultOutputInterval,
		},
		Log: LogConfig{
			Level:  DefaultLogLevel,
			Format: DefaultLogFormat,
			Output: DefaultLogOutput,
		},
		Web: WebConfig{
			Port:    DefaultWebPort,
			CORS:    DefaultCORSOrigins,
			Enabled: false,
			APIPath: DefaultAPIPath,
		},
		Parser: ParserConfig{
			Method:     DefaultParsingMethod,
			Concurrent: true,
			RegexLimit: DefaultRegexLimit,
		},
		Output: OutputConfig{
			Format:      DefaultOutputFormat,
			Colors:      true,
			Indentation: DefaultIndentation,
		},
	}
}

// LoadFromFile loads configuration from a config file
// Supports JSON and YAML formats
func (c *Config) LoadFromFile() error {
	// TODO: Implement config file loading
	// Look for: goperf.json, goperf.yaml, .goperf.json, .goperf.yaml
	return nil
}

// LoadFromEnv loads configuration from environment variables
func (c *Config) LoadFromEnv() error {
	// HTTP configuration
	if timeout := os.Getenv("GOPERF_HTTP_TIMEOUT"); timeout != "" {
		if d, err := time.ParseDuration(timeout); err == nil {
			c.HTTP.Timeout = d
		}
	}

	if maxConn := os.Getenv("GOPERF_HTTP_MAX_CONNECTIONS"); maxConn != "" {
		if n, err := strconv.Atoi(maxConn); err == nil {
			c.HTTP.MaxConnections = n
		}
	}

	if userAgent := os.Getenv("GOPERF_USER_AGENT"); userAgent != "" {
		c.HTTP.UserAgent = userAgent
	}

	// Test configuration
	if users := os.Getenv("GOPERF_DEFAULT_USERS"); users != "" {
		if n, err := strconv.Atoi(users); err == nil {
			c.Test.DefaultUsers = n
		}
	}

	if url := os.Getenv("GOPERF_DEFAULT_URL"); url != "" {
		c.Test.DefaultURL = url
	}

	// Web configuration
	if port := os.Getenv("GOPERF_WEB_PORT"); port != "" {
		if n, err := strconv.Atoi(port); err == nil {
			c.Web.Port = n
		}
	}

	// Log configuration
	if level := os.Getenv("GOPERF_LOG_LEVEL"); level != "" {
		c.Log.Level = level
	}

	return nil
}

// LoadFromFlags loads configuration from command line flags
func (c *Config) LoadFromFlags() error {
	// Define flags
	users := flag.Int("users", c.Test.DefaultUsers, "Number of concurrent users/connections")
	url := flag.String("url", c.Test.DefaultURL, "URL to test")
	seconds := flag.Int("sec", int(c.Test.DefaultDuration.Seconds()), "Test duration in seconds")
	web := flag.Bool("web", c.Web.Enabled, "Run as a webserver")
	port := flag.Int("port", c.Web.Port, "Web server port")
	userAgent := flag.String("useragent", c.HTTP.UserAgent, "User agent string")
	outputFile := flag.String("output", c.Test.OutputFile, "Output file path")

	// Parse flags
	flag.Parse()

	// Apply flag values
	c.Test.DefaultUsers = *users
	c.Test.DefaultURL = *url
	c.Test.DefaultDuration = time.Duration(*seconds) * time.Second
	c.Web.Enabled = *web
	c.Web.Port = *port
	c.HTTP.UserAgent = *userAgent
	c.Test.OutputFile = *outputFile

	return nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.HTTP.Timeout <= 0 {
		return fmt.Errorf("HTTP timeout must be positive")
	}

	if c.HTTP.MaxConnections <= 0 {
		return fmt.Errorf("max connections must be positive")
	}

	if c.Test.DefaultUsers <= 0 {
		return fmt.Errorf("default users must be positive")
	}

	if c.Web.Port < MinPortNumber || c.Web.Port > MaxPortNumber {
		return fmt.Errorf("web port must be between %d and %d", MinPortNumber, MaxPortNumber)
	}

	return nil
}
