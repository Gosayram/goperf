package interfaces

// AssetParser defines the contract for extracting assets from HTML content
// This replaces the current httputils package direct function calls
type AssetParser interface {
	// ParseAssets extracts all assets (JS, CSS, images) from HTML content
	ParseAssets(body string) (*Assets, error)

	// ParseJS extracts JavaScript file URLs from HTML content
	ParseJS(body string) ([]string, error)

	// ParseCSS extracts CSS file URLs from HTML content
	ParseCSS(body string) ([]string, error)

	// ParseImages extracts image URLs from HTML content
	ParseImages(body string) ([]string, error)

	// SetParsingMethod allows switching between regex and DOM parsing
	SetParsingMethod(method ParsingMethod)
}

// ParsingMethod represents the method used for parsing HTML content
type ParsingMethod int

const (
	// ParsingMethodRegex indicates regex-based parsing for asset extraction
	ParsingMethodRegex ParsingMethod = iota
	// ParsingMethodDOM indicates DOM-based parsing using goquery
	ParsingMethodDOM
	// ParsingMethodMixed indicates a combination of regex and DOM parsing
	ParsingMethodMixed
)

// Assets represents all extracted assets from a page
type Assets struct {
	JavaScript []string `json:"javascript"`
	CSS        []string `json:"css"`
	Images     []string `json:"images"`
	Total      int      `json:"total"`
}

// String returns a string representation of parsing method
func (p ParsingMethod) String() string {
	switch p {
	case ParsingMethodRegex:
		return "regex"
	case ParsingMethodDOM:
		return "dom"
	case ParsingMethodMixed:
		return "mixed"
	default:
		return "unknown"
	}
}
