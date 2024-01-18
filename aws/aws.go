package aws

import (
	"regexp"
	"strings"
)

const (
	PublicKeyNameSuffix  = "_pub.pem"
	PrivateKeyNameSuffix = "_priv.pem"
)

const ParameterStoreNamingRequirementsString = `
Invalid parameter store name: '%s'. 

Ensure the name follows AWS Parameter Store naming conventions: 
it must start with a letter or a number, can only include 
alphanumeric characters, hyphens (-), underscores (_), and
forward slashes (/), and must not exceed 2048 characters in length.
`

func MakePrivateKeyName(keyName string) string {
	return keyName + PrivateKeyNameSuffix
}

func MakePublicKeyName(keyName string) string {
	return keyName + PublicKeyNameSuffix
}

// IsValidParameterStoreName checks if a string is a valid AWS Parameter Store name.
func IsValidParameterStoreName(name string) bool {
	// AWS Parameter Store names can contain letters, numbers, hyphens, and underscores.
	// They must start with a letter or number and can't exceed 2048 characters.
	// We'll use a regular expression to validate the name.
	// Adjust the regular expression pattern as needed.

	// Regular expression pattern for a valid AWS Parameter Store name.
	pattern := `^[a-zA-Z0-9][a-zA-Z0-9_\-/]*$`

	// Compile the regular expression.
	regex := regexp.MustCompile(pattern)

	// Check if the name matches the pattern.
	if !regex.MatchString(name) {
		return false
	}

	// Check if the name doesn't exceed 2048 characters.
	return len(name) <= 2048
}

func GetFilenameFromParameterStorePath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
