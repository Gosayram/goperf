package request

const (
	// HTTPStatusOK represents the HTTP 200 status code for successful requests
	HTTPStatusOK = 200
	// HTTPStatusNotFound represents the HTTP 404 status code for not found resources
	HTTPStatusNotFound = 404
	// HTTPStatusConnectionError represents a connection error status code
	HTTPStatusConnectionError = -100 // Connection error status

	// HTTPProtocolLength specifies the length of the HTTP protocol prefix
	HTTPProtocolLength = 4 // Length of "http" string

	// AssetTypesCount specifies the number of different asset types (JS, CSS, images)
	AssetTypesCount = 3 // Number of asset types: JS, CSS, IMG

	// FirstPathElement specifies the index of the first path element in URL parsing
	FirstPathElement = 0 // Index of first character in path
	// SecondPathElement specifies the index of the second path element in URL parsing
	SecondPathElement = 1 // Index of second character in path

	// DefaultJSONIndent specifies the default indentation for JSON formatting
	DefaultJSONIndent = "  " // Default JSON indentation
	// DefaultEmptyString specifies the default empty string value for string fields
	DefaultEmptyString = "" // Default empty string value
	// DefaultTimeout specifies the default timeout duration in seconds for HTTP requests
	DefaultTimeout = 5 // Default timeout in seconds
	// DefaultBufferSize specifies the default buffer size in bytes for data operations
	DefaultBufferSize = 1024 // Default buffer size in bytes

	// UserAgentHeader specifies the HTTP User-Agent header name
	UserAgentHeader = "User-Agent"
	// CookieHeader specifies the HTTP Cookie header name
	CookieHeader = "cookie"

	// ErrorRequestFailed specifies the default error message for failed requests
	ErrorRequestFailed = "Request failed"

	// HTTPScheme specifies the HTTP protocol scheme for URLs
	HTTPScheme = "http"
	// HTTPSScheme specifies the HTTPS protocol scheme for secure URLs
	HTTPSScheme = "https"

	// PathSeparator specifies the path separator character for URLs
	PathSeparator = "/"
)
