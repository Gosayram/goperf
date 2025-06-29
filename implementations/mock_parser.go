package implementations

import (
	"github.com/Gosayram/goperf/interfaces"
)

// MockAssetParser is a simple implementation for testing
type MockAssetParser struct {
	method interfaces.ParsingMethod
}

// NewMockAssetParser creates a new mock asset parser
func NewMockAssetParser() *MockAssetParser {
	return &MockAssetParser{
		method: interfaces.ParsingMethodDOM,
	}
}

// ParseAssets implements interfaces.AssetParser
func (p *MockAssetParser) ParseAssets(_ string) (*interfaces.Assets, error) {
	// Mock asset parsing - return some fake assets
	return &interfaces.Assets{
		JavaScript: []string{
			"/assets/app.js",
			"/assets/vendor.js",
		},
		CSS: []string{
			"/assets/style.css",
			"/assets/theme.css",
		},
		Images: []string{
			"/images/logo.png",
			"/images/background.jpg",
		},
		Total: MockAssetCount,
	}, nil
}

// ParseJS implements interfaces.AssetParser
func (p *MockAssetParser) ParseJS(_ string) ([]string, error) {
	return []string{
		"/assets/app.js",
		"/assets/vendor.js",
	}, nil
}

// ParseCSS implements interfaces.AssetParser
func (p *MockAssetParser) ParseCSS(_ string) ([]string, error) {
	return []string{
		"/assets/style.css",
		"/assets/theme.css",
	}, nil
}

// ParseImages implements interfaces.AssetParser
func (p *MockAssetParser) ParseImages(_ string) ([]string, error) {
	return []string{
		"/images/logo.png",
		"/images/background.jpg",
	}, nil
}

// SetParsingMethod implements interfaces.AssetParser
func (p *MockAssetParser) SetParsingMethod(method interfaces.ParsingMethod) {
	p.method = method
}
