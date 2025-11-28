package utils

import (
	"strings"
)

const (
	SEXUAL = "{{性的}}"
)

func judgeIsSexual(s string) bool {
	return strings.Contains(s, SEXUAL) 
}