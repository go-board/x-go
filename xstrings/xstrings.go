package xstrings

import (
	"github.com/huandu/xstrings"
)

// Reverse reverse str rune by rune
func Reverse(str string) string {
	return xstrings.Reverse(str)
}

// CamelCast transform a string to camel case
func CamelCase(str string) string {
	return xstrings.ToCamelCase(str)
}

// SnakeCast transform a string to snake case
func SnakeCase(str string) string {
	return xstrings.ToSnakeCase(str)
}

// KebabCast transform a string to kebab case
func KebabCase(str string) string {
	return xstrings.ToKebabCase(str)
}

// Delete delete all `pattern` in `str`
func Delete(str string, pattern string) string {
	return xstrings.Delete(str, pattern)
}

// Count count `pattern` appear times in `str`
func Count(str string, pattern string) int {
	return xstrings.Count(str, pattern)
}
