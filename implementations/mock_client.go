package implementations

import (
	"context"
	"fmt"
	"time"

	"github.com/Gosayram/goperf/interfaces"
)

// MockHTTPClient is a simple implementation for testing the new architecture
type MockHTTPClient struct {
	timeout        time.Duration
	userAgent      string
	maxConnections int
}

// NewMockHTTPClient creates a new mock HTTP client
func NewMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{
		timeout:        MockHTTPTimeout,
		userAgent:      "goperf",
		maxConnections: MockMaxConnections,
	}
}

// Fetch implements interfaces.HTTPClient
func (c *MockHTTPClient) Fetch(_ context.Context, req *interfaces.Request) (*interfaces.Response, error) {
	// Simulate HTTP request
	start := time.Now()

	// Simulate some processing time
	time.Sleep(MockSleepDuration * time.Millisecond)

	duration := time.Since(start)

	return &interfaces.Response{
		URL:        req.URL,
		StatusCode: MockHTTPStatusOK,
		Headers:    make(map[string][]string),
		Body:       fmt.Sprintf("Mock response for %s", req.URL),
		Size:       MockResponseSize,
		Duration:   duration,
		Error:      nil,
	}, nil
}

// FetchBatch implements interfaces.HTTPClient
func (c *MockHTTPClient) FetchBatch(ctx context.Context, req *interfaces.Request) (*interfaces.BatchResponse, error) {
	// Fetch base response
	baseResp, err := c.Fetch(ctx, req)
	if err != nil {
		return nil, err
	}

	// Mock some assets
	assets := []*interfaces.Response{
		{
			URL:        req.URL + "/assets/style.css",
			StatusCode: MockHTTPStatusOK,
			Size:       MockMediumAssetSize,
			Duration:   MockMediumLatency * time.Millisecond,
		},
		{
			URL:        req.URL + "/assets/script.js",
			StatusCode: MockHTTPStatusOK,
			Size:       MockSmallAssetSize,
			Duration:   MockShortLatency * time.Millisecond,
		},
	}

	totalSize := baseResp.Size
	for _, asset := range assets {
		totalSize += asset.Size
	}

	return &interfaces.BatchResponse{
		BaseResponse: baseResp,
		Assets:       assets,
		TotalTime:    baseResp.Duration + 8*time.Millisecond, // Combined asset time
		TotalSize:    totalSize,
	}, nil
}

// SetTimeout implements interfaces.HTTPClient
func (c *MockHTTPClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// SetUserAgent implements interfaces.HTTPClient
func (c *MockHTTPClient) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

// SetMaxConnections implements interfaces.HTTPClient
func (c *MockHTTPClient) SetMaxConnections(maxConns int) {
	c.maxConnections = maxConns
}
