// Package implementations provides mock implementations of core interfaces
// for testing and demonstration purposes. These implementations can be used
// to simulate various scenarios without making actual HTTP requests.
package implementations

import "time"

const (
	// MockHTTPTimeout specifies the default timeout for mock HTTP requests
	MockHTTPTimeout = 5 * time.Second // Mock HTTP timeout
	// MockMaxConnections specifies the maximum number of mock HTTP connections
	MockMaxConnections = 100 // Mock maximum connections
	// MockRetryAttempts specifies the number of retry attempts for mock requests
	MockRetryAttempts = 3 // Mock retry attempts

	// MockHTTPStatusOK specifies the HTTP status code for successful mock responses
	MockHTTPStatusOK = 200 // Mock successful HTTP status
	// MockHTTPStatusNotFound specifies the HTTP 404 status for mock not found responses
	MockHTTPStatusNotFound = 404 // Mock not found HTTP status
	// MockHTTPStatusError specifies the HTTP 500 status for mock server error responses
	MockHTTPStatusError = 500 // Mock server error HTTP status

	// MockShortLatency specifies a short latency value for mock responses
	MockShortLatency = 10 * time.Millisecond // Mock short response latency
	// MockMediumLatency specifies a medium latency value for mock responses
	MockMediumLatency = 50 * time.Millisecond // Mock medium response latency

	// MockTotalRequests specifies the total number of mock requests
	MockTotalRequests = 100 // Mock total request count
	// MockSuccessRequests specifies the number of successful mock requests
	MockSuccessRequests = 95 // Mock successful request count
	// MockFailedRequests specifies the number of failed mock requests
	MockFailedRequests = 5 // Mock failed request count
	// MockAvgLatency specifies the average latency value for mock responses
	MockAvgLatency = 25 * time.Millisecond // Mock average latency
	// MockMinLatency specifies the minimum latency value for mock responses
	MockMinLatency = 10 * time.Millisecond // Mock minimum latency
	// MockMaxLatency specifies the maximum latency value for mock responses
	MockMaxLatency = 100 * time.Millisecond // Mock maximum latency

	// MockAssetCount specifies the number of mock assets to generate
	MockAssetCount = 10 // Mock asset count for responses
	// MockJSAssetCount specifies the number of mock JavaScript assets
	MockJSAssetCount = 2 // Mock JavaScript asset count
	// MockCSSAssetCount specifies the number of mock CSS assets
	MockCSSAssetCount = 1 // Mock CSS asset count
	// MockImageAssetCount specifies the number of mock image assets
	MockImageAssetCount = 3 // Mock image asset count

	// MockSleepDuration specifies the sleep duration for mock operations
	MockSleepDuration = 100 * time.Millisecond // Mock sleep duration
	// MockResponseSize specifies the mock response size in bytes for mock data
	MockResponseSize = 100 // Mock response size in bytes
	// MockSmallAssetSize specifies the size of small mock assets in bytes
	MockSmallAssetSize = 30 // Mock small asset size in bytes
	// MockMediumAssetSize specifies the size of medium mock assets in bytes
	MockMediumAssetSize = 50 // Mock medium asset size in bytes
	// MockLargeAssetSize specifies the size of large mock assets in bytes
	MockLargeAssetSize = 100 // Mock large asset size in bytes
	// MockThroughput specifies the mock throughput in requests per second
	MockThroughput = 12.5 // Mock throughput in requests per second
	// MockTotalBytes specifies the mock total bytes transferred during testing
	MockTotalBytes = 2500 // Mock total bytes transferred
)
