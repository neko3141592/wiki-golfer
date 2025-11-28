package utils

import (
	"regexp"
	"strings"
)

func extractTitle(articleLines []string) string {
	for _, line := range articleLines {
		if strings.Contains(line, "<title>") {
			start := strings.Index(line, "<title>") + len("<title>")
			end := strings.Index(line, "</title>")
			if start >= 0 && end > start {
				return line[start:end]
			}
		}
	}
	return ""
}

func extractLinks(articleLines []string) []string {
	text := strings.Join(articleLines, "")
	re := regexp.MustCompile(`\[\[([^\[\]|]+)`)
	matches := re.FindAllStringSubmatch(text, -1)
	var links []string
	for _, m := range matches {
		links = append(links, m[1])
	}
	return links
}

func extractID(articleLines []string) string {
	for _, line := range articleLines {
		if strings.Contains(line, "<id>") {
			start := strings.Index(line, "<id>") + len("<id>")
			end := strings.Index(line, "</id>")
			if start >= 0 && end > start {
				return line[start:end]
			}
		}
	}
	return ""
}