// Package httputils provides utilities for parsing HTML, CSS, and JavaScript assets
// from web pages. It includes functions for extracting asset URLs using regex patterns
// and the goquery library for fast DOM parsing.
package httputils

const (
	// AssetTypesCount specifies the number of different asset types (JS, CSS, images)
	AssetTypesCount = 3 // Number of asset types processed
	// RegexMatchLimit specifies the limit for regex matches during asset extraction
	RegexMatchLimit = -10 // Unlimited matches for regex processing

	// HTTPProtocolLength specifies the length of HTTP protocol prefix in URLs
	HTTPProtocolLength = 4 // Length of "http" string

	// ScriptTag specifies the HTML script tag name for JavaScript assets
	ScriptTag = "script" // HTML script tag
	// ImageTag specifies the HTML img tag name for image assets
	ImageTag = "img" // HTML img tag
	// LinkTag specifies the HTML link tag name for CSS and other linked assets
	LinkTag = "link" // HTML link tag
	// SrcAttribute specifies the HTML src attribute name for asset URLs
	SrcAttribute = "src" // HTML src attribute
	// HrefAttribute specifies the HTML href attribute name for link URLs
	HrefAttribute = "href" // HTML href attribute

	// ScriptSrcPattern specifies the regex pattern for extracting script source URLs
	ScriptSrcPattern = `<script.*?src=["'\''](.*?)["'\''].*?>`
	// LinkHrefPattern specifies the regex pattern for extracting link href URLs
	LinkHrefPattern = `<link.*?href=["'\''](.*?)["'\''].*?>`
	// ImageSrcPattern specifies the regex pattern for extracting img src URLs
	ImageSrcPattern = `<img.*?src=["'\''](.*?)["'\''].*?>`
	// BackgroundImagePattern specifies the regex pattern for extracting background image URLs from CSS
	BackgroundImagePattern = `background-image\s*:\s*url\(["'\'']?(.*?)["'\'']?\)`
)
