package main

import (
	// "fmt"
	"strings"
)

func Camel(arg string) string {
	var sb strings.Builder
	for _, ch := range arg {
		// Looking for Upper case to stop and replace by _
		if ch >= 'A' && ch <= 'Z' {
			sb.WriteRune('_')
		}
		// Write that whole sentence
		sb.WriteRune(ch)
	}

	// Return as lower
	return strings.ToLower((sb.String()))
}
