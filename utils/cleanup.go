package utils

import "strings"

func CleanupString(s string) string {
	return strings.ReplaceAll(s, "'", "")
}
