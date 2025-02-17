package utils

import "strings"

func NormalizeToken(token string) string {
	return strings.ReplaceAll(token, "-", "")
}
