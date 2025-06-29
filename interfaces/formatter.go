package interfaces

// OutputFormatter defines the contract for formatting test results
// This replaces the scattered JSON formatting throughout the codebase
type OutputFormatter interface {
	// FormatJSON formats data as JSON with proper indentation
	FormatJSON(data interface{}) ([]byte, error)

	// FormatText formats data as human-readable text with colors
	FormatText(data interface{}) (string, error)

	// FormatCSV formats data as CSV for spreadsheet import
	FormatCSV(data interface{}) ([]byte, error)

	// FormatHTML formats data as HTML report
	FormatHTML(data interface{}) ([]byte, error)

	// SetIndentation configures JSON indentation
	SetIndentation(indent string)

	// SetColors enables/disables colored output for text format
	SetColors(enabled bool)
}

// OutputFormat represents the format for output data
type OutputFormat int

const (
	// OutputFormatJSON represents JSON output format
	OutputFormatJSON OutputFormat = iota
	// OutputFormatText represents text output format
	OutputFormatText
	// OutputFormatCSV represents CSV output format
	OutputFormatCSV
	// OutputFormatHTML represents HTML output format
	OutputFormatHTML
	// UnknownFormat represents an unknown or unsupported output format
	UnknownFormat = "unknown"
)

// String returns the string representation of the output format
func (f OutputFormat) String() string {
	switch f {
	case OutputFormatJSON:
		return "json"
	case OutputFormatText:
		return "text"
	case OutputFormatCSV:
		return "csv"
	case OutputFormatHTML:
		return "html"
	default:
		return UnknownFormat
	}
}
