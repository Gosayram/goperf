// Package interfaces defines the core interfaces for the GoPerf load testing tool.
// These interfaces provide contracts for HTTP clients, asset parsers, metrics collectors,
// and output formatters, enabling clean architecture and dependency injection.
package interfaces

import (
	"context"
	"time"
)

// HTTPClient defines the contract for making HTTP requests
// This replaces the direct http.Client usage throughout the codebase
type HTTPClient interface {
	// Fetch performs a single HTTP request
	Fetch(ctx context.Context, req *Request) (*Response, error)

	// FetchBatch performs multiple HTTP requests concurrently
	// This replaces the current FetchAll functionality
	FetchBatch(ctx context.Context, req *Request) (*BatchResponse, error)

	// SetTimeout configures request timeout
	SetTimeout(timeout time.Duration)

	// SetUserAgent configures the User-Agent header
	SetUserAgent(userAgent string)

	// SetMaxConnections configures connection pooling
	SetMaxConnections(maxConns int)
}

// Request represents a unified HTTP request structure
// This replaces FetchInput and other similar structs
type Request struct {
	URL           string            `json:"url"`
	Method        string            `json:"method"`
	Headers       map[string]string `json:"headers"`
	Cookies       string            `json:"cookies"`
	UserAgent     string            `json:"user_agent"`
	Timeout       time.Duration     `json:"timeout"`
	ReturnContent bool              `json:"return_content"`
}

// Response represents a unified HTTP response structure
// This replaces FetchResponse and other similar structs
type Response struct {
	URL        string              `json:"url"`
	StatusCode int                 `json:"status_code"`
	Headers    map[string][]string `json:"headers"`
	Body       string              `json:"body"`
	Size       int                 `json:"size"`
	Duration   time.Duration       `json:"duration"`
	Error      error               `json:"error,omitempty"`
}

// BatchResponse represents the result of fetching a page with all assets
// This replaces FetchAllResponse struct
type BatchResponse struct {
	BaseResponse *Response     `json:"base_response"`
	Assets       []*Response   `json:"assets"`
	TotalTime    time.Duration `json:"total_time"`
	TotalSize    int           `json:"total_size"`
}
