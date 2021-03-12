package helpers

import "unicode"

// HasMixedCase verifies if a string has both upper and lower case characters.
func HasMixedCase(str string) bool {
	var hasLower, hasUpper bool
	hasLower = false
	hasUpper = false
	for _, character := range str {
		if unicode.IsLetter(character) {
			if hasLower == false {
				hasLower = unicode.IsLower(character)
			}
			if hasUpper == false {
				hasUpper = unicode.IsUpper(character)
			}
		}
		if hasLower && hasUpper {
			return true
		}
	}
	return false
}

// HasSpecialChars checks to see if a string has any special (printable) characters.
func HasSpecialChar(str string) bool {
	var specialChar = "!@#$%^&*"
	for _, character := range str {
		for _, spcCharacter := range specialChar {
			if character == spcCharacter { return true}
		}
	}
	return false
}

// HasAlphaNum checks to see if a string has both letters and numbers.
func HasAlphaNum(str string) bool {
	var hasAlpha, hasNumber bool
	hasAlpha = false
	hasNumber = false

	for _, character := range str {
		if unicode.IsLetter(character) {
			hasAlpha = true
		}
		if unicode.IsNumber(character) {
			hasNumber = true
		}
		if (hasAlpha && hasNumber) {
			return true
		}
	}
	return false
}