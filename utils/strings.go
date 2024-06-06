package utils

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CutStringFromMatch cuts the string from the first match of the pattern
func CutStringFromMatch(input, pattern string) string {
	parts := strings.Split(input, pattern)
	if len(parts) < 2 {
		return input
	}
	return strings.TrimSpace(parts[1])
}

// ToCamel converts string to camelCase
func ToCamel(s string) string {
	words := SplitIntoWords(s)
	caser := cases.Title(language.English)
	for i := 1; i < len(words); i++ {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}

// ToKebab converts string to kebab-case
func ToKebab(s string) string {
	return strings.Join(SplitIntoWords(s), "-")
}

// ToTitle converts string to Title Case
func ToTitle(s string) string {
	words := SplitIntoWords(s)
	caser := cases.Title(language.English)
	for i := range words {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, " ")
}

// ToSnake converts string to snake_case
func ToSnake(s string) string {
	return strings.Join(SplitIntoWords(s), "_")
}

// ToPascal converts string to PascalCase
func ToPascal(s string) string {
	words := SplitIntoWords(s)
	caser := cases.Title(language.English)
	for i := range words {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}

// ToConstant converts string to CONSTANT_CASE
func ToConstant(s string) string {
	return strings.ToUpper(strings.Join(SplitIntoWords(s), "_"))
}

// SplitIntoWords splits the string into words based on non-alphanumeric characters
func SplitIntoWords(s string) []string {
	return regexp.MustCompile("[^a-zA-Z0-9]+").Split(strings.ToLower(s), -1)
}
