package helpers

import (
	"fmt"
	"strings"
)

// SSV -> Semi-colon delimited values
// func ParseStringFromSSV(input string) map[string]string {
func ParseStringFromSSV(input string) ([][]string, error) {
	var result [][]string

	// Custom splitting function that respects quotes
	pairs := splitWithQuotes(input)

	for _, pair := range pairs {
		// Trim whitespace
		pair = strings.TrimSpace(pair)

		// Split on first '=' sign
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid key-value pair: %s", pair)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
		}

		// Unescape semicolons
		value = strings.ReplaceAll(value, `\;`, `;`)

		result = append(result, []string{key, value})
	}

	return result, nil
}

// splitWithQuotes custom function to split on semicolons while respecting quotes
func splitWithQuotes(s string) []string {
	var result []string
	var current strings.Builder
	var inQuotes bool
	var escaped bool

	for _, r := range s {
		switch {
		case escaped:
			// If previous char was a backslash, add this char regardless
			current.WriteRune(r)
			escaped = false
		case r == '\\':
			// Mark next char as escaped
			escaped = true
		case r == '"':
			// Toggle quote state
			inQuotes = !inQuotes
			current.WriteRune(r)
		case r == ';' && !inQuotes:
			// Semicolon outside quotes means split
			result = append(result, current.String())
			current.Reset()
		default:
			// Normal character
			current.WriteRune(r)
		}
	}

	// Add the last segment
	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}
