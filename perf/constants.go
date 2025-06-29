// Package perf provides performance testing functionality for concurrent website load testing.
// It includes structures for managing test iterations, collecting metrics, and formatting
// results in various output formats including JSON and console output.
package perf

const (
	// DefaultSleepDuration specifies the default sleep duration between operations
	DefaultSleepDuration = 1000 // milliseconds
	// PercentageBase specifies the base value for percentage calculations
	PercentageBase = 100.0 // Base for percentage calculations
	// FloatPrecision specifies the precision level for floating-point formatting
	FloatPrecision = 5 // Precision for float formatting
	// Radix64 specifies the base 64 radix for string conversion operations
	Radix64 = 64 // Base 64 for string conversion

	// DefaultJSONIndent specifies the default indentation for JSON formatting
	DefaultJSONIndent = "  " // 2 spaces for JSON formatting

	// AssetTypesCount specifies the number of different asset types processed
	AssetTypesCount = 3 // JS, CSS, IMG

	// FloatFormatChar specifies the format character for float conversion
	FloatFormatChar = 'g' // Float format character for strconv
)
