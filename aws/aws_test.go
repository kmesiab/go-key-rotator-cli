package aws

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFilenameFromParameterStorePath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Valid Paths/Names
		{"/myapp/secrets/api_key", "api_key"},
		{"password", "password"},

		// Invalid Paths/Names
		{"", ""},                         // Empty String
		{"/", ""},                        // Leading Slash Only
		{"myparam/", ""},                 // Trailing Slash Only
		{"/myapp//secrets", "secrets"},   // Multiple Consecutive Slashes
		{"/my param/$secret", "$secret"}, // Special Characters and Spaces
		{"/my_param/@token", "@token"},   // Non-alphanumeric Characters
		{"/my app/secrets/key", "key"},   // Path with Spaces
		{"/myapp/你好/こんにちは", "こんにちは"}, // Path with Unicode Characters
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := GetFilenameFromParameterStorePath(test.input)
			if result != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, result)
			}
		})
	}
}

func TestIsValidParameterStoreName(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"validName", true},
		{"myapp/secrets/api_key", true},
		{"123", true},       // Invalid: starts with a number
		{"", false},         // Invalid: empty name
		{longName, false},   // Invalid: exceeds 2048 characters
		{longishName, true}, // Valid: almost 2048 characters
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidParameterStoreName(test.name)
			assert.Equal(t, test.expected, result)
		})
	}
}

// Define a long name to test the 2048-character limit max.
var longName = "a" + strings.Repeat("b", 2050)

// Define a long name to test the 2048-character limit.
var longishName = "a" + strings.Repeat("b", 2045)

func TestMakePrivateKeyName(t *testing.T) {
	tests := []struct {
		keyName  string
		expected string
	}{
		{"myKey", "myKey" + PrivateKeyNameSuffix},
		{"testKey", "testKey" + PrivateKeyNameSuffix},
		{"", PrivateKeyNameSuffix},
	}
	for _, test := range tests {
		t.Run(test.keyName, func(t *testing.T) {
			result := MakePrivateKeyName(test.keyName)
			assert.Equal(t, test.expected, result,
				"Output should match expected private key name")
		})
	}
}

func TestMakePublicKeyName(t *testing.T) {
	tests := []struct {
		keyName  string
		expected string
	}{
		{"myKey", "myKey" + PublicKeyNameSuffix},
		{"testKey", "testKey" + PublicKeyNameSuffix},
		{"", PublicKeyNameSuffix},
	}

	for _, test := range tests {
		t.Run(test.keyName, func(t *testing.T) {
			result := MakePublicKeyName(test.keyName)
			assert.Equal(t, test.expected, result,
				"Output should match expected public key name")
		})
	}
}
