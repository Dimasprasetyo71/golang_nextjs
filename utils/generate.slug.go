package utils

import (
	"strings"
)

func GenerateSlug(s string) string {
	slug := strings.ReplaceAll(s, " ", "-")

	return slug
}

func GenerateUUID() string {
	return RandomString(36)
}
