package core

import "time"

const (
	// DefaultHTTPTimeout specifies the default timeout for HTTP requests
	DefaultHTTPTimeout = 30 * time.Second // Default HTTP request timeout
	// DefaultMaxConnections specifies the maximum number of concurrent HTTP connections
	DefaultMaxConnections = 100 // Default maximum HTTP connections
	// DefaultRetryAttempts specifies the default number of retry attempts for failed requests
	DefaultRetryAttempts = 3 // Default number of retry attempts
	// DefaultUserAgent specifies the default User-Agent header for HTTP requests
	DefaultUserAgent = "goperf" // Default User-Agent header

	// DefaultUsers specifies the default number of concurrent users for load testing
	DefaultUsers = 1 // Default number of concurrent users
	// DefaultTestDuration specifies the default duration for performance tests
	DefaultTestDuration = 2 * time.Second // Default test duration
	// DefaultURL specifies the default URL for load testing operations
	DefaultURL = "https://httpbin.org/get" // Default test URL
	// DefaultOutputFile specifies the default output file path for test results
	DefaultOutputFile = "./output.json" // Default output file path
	// DefaultIterations specifies the default number of test iterations
	DefaultIterations = 1000 // Default number of iterations
	// DefaultOutputInterval specifies the default interval for output reporting in seconds
	DefaultOutputInterval = 5 // Default output interval

	// DefaultWebPort specifies the default port for web server mode
	DefaultWebPort = 8080 // Default web server port
	// MinPortNumber specifies the minimum valid port number for network services
	MinPortNumber = 1 // Minimum valid port number
	// MaxPortNumber specifies the maximum valid port number for network services
	MaxPortNumber = 65535 // Maximum valid port number

	// DefaultRegexLimit specifies the default limit for regex matches during parsing
	DefaultRegexLimit = -10 // Default regex match limit (-1 means unlimited)

	// MockHTTPStatusOK represents a successful HTTP status code for mock responses
	MockHTTPStatusOK = 200 // Mock HTTP 200 status
	// DefaultSuccessRate specifies the default success rate percentage
	DefaultSuccessRate = 100.0 // Default success rate percentage
	// MockResponseSize specifies the mock response size in bytes for testing
	MockResponseSize = 100 // Mock response size in bytes
	// MockLatency specifies the mock latency in milliseconds for testing
	MockLatency = 10 // Mock latency in milliseconds
	// MockTotalRequests specifies the mock total requests count for testing
	MockTotalRequests = 25 // Mock total requests count
	// MockSuccessRequests specifies the mock successful requests count for testing
	MockSuccessRequests = 25 // Mock successful requests count
	// MockAvgLatency specifies the mock average latency in milliseconds for testing
	MockAvgLatency = 242 // Mock average latency in milliseconds
	// MockMinLatency specifies the minimum latency for mock responses
	MockMinLatency = 180 // Mock minimum latency in milliseconds
	// MockMaxLatency specifies the maximum latency for mock responses
	MockMaxLatency = 350 // Mock maximum latency in milliseconds
	// MockThroughput specifies the mock throughput in requests per second for testing
	MockThroughput = 12.5 // Mock throughput (requests per second)
	// MockTotalBytes specifies the mock total bytes transferred for testing
	MockTotalBytes = 2500 // Mock total bytes transferred
	// MockSmallAssetSize specifies the mock small asset size for testing
	MockSmallAssetSize = 30 // Mock small asset size
	// MockMediumAssetSize specifies the mock medium asset size for testing
	MockMediumAssetSize = 50 // Mock medium asset size
	// MockShortLatency specifies a short latency value for mock responses
	MockShortLatency = 3 // Mock short latency
	// MockMediumLatency specifies a medium latency value for mock responses
	MockMediumLatency = 5 // Mock medium latency
	// MockAssetCount specifies the number of mock assets to generate
	MockAssetCount = 6 // Mock asset count
	// MockSleepDuration specifies the mock sleep duration in milliseconds for testing
	MockSleepDuration = 10 // Mock sleep duration in milliseconds

	// DefaultLogLevel specifies the default logging level
	DefaultLogLevel = "info" // Default log level
	// DefaultLogFormat specifies the default log format for output
	DefaultLogFormat = "text" // Default log format
	// DefaultLogOutput specifies the default log output destination
	DefaultLogOutput = "stdout" // Default log output
	// DefaultParsingMethod specifies the default HTML parsing method
	DefaultParsingMethod = "dom" // Default HTML parsing method
	// DefaultOutputFormat specifies the default output format for results
	DefaultOutputFormat = "text" // Default output format
	// DefaultIndentation specifies the default JSON indentation string
	DefaultIndentation = "    " // Default JSON indentation
	// DefaultAPIPath specifies the default API path for web server mode
	DefaultAPIPath = "/api/" // Default API path

	// TotalTimeInMilliseconds specifies the total time for batch response processing
	TotalTimeInMilliseconds = 15 // Total time for batch response in milliseconds
	// ContainerLineLimit specifies the maximum characters per line in output
	ContainerLineLimit = 120 // Maximum characters per line
)

var (
	// DefaultCORSOrigins specifies the default CORS origins for web server mode
	DefaultCORSOrigins = []string{"*", "http://localhost:8080"}
)
