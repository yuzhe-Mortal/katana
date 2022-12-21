package utils

import (
	"regexp"
)

var (
	// pageBodyRegex extracts endpoints from page body
	pageBodyRegex = regexp.MustCompile(
		`((?:(?:[\.]{1,2}/[A-Za-z0-9-_/\\?&@\.?=]+)|https?://[A-Za-z0-9_\-\.]+([\.]{0,2})?\/[A-Za-z0-9-_/\\?&@\.?=]+|(/[A-Za-z0-9-_/\\?&@\.]+\.(aspx?|action|cfm|cgi|do|pl|css|x?html?|js(p|on)?|pdf|php5?|py|rss))))`,
	)
	// relativeEndpointsRegex is the regex to find endpoints in js files.
	relativeEndpointsRegex = regexp.MustCompile(
		`(?:"|'|\s)(((https?://[A-Za-z0-9_\-\.]+(:\d{1,5})?)+([\.]{1,2})?/[A-Za-z0-9/\-_\.\\]+([\?|#][^"']+)?)|((\.{1,2}/)?[a-zA-Z0-9\-_/\\]+\.(aspx?|js(on|p)?|html|php5?|html|action|do)([\?|#][^"']+)?)|((\.{0,2}/)[a-zA-Z0-9\-_/\\]+(/|\\)[a-zA-Z0-9\-_]{3,}([\?|#][^"|']+)?))(?:"|'|\s)`,
	)
)

// ExtractBodyEndpoints extracts body endpoints from a data item
func ExtractBodyEndpoints(data string) []string {
	matches := []string{}
	unique := make(map[string]struct{})

	relativeMatches := pageBodyRegex.FindAllStringSubmatch(data, -1)
	for _, match := range relativeMatches {
		if len(match) < 2 {
			continue
		}
		if _, ok := unique[match[1]]; ok {
			continue
		}
		unique[match[1]] = struct{}{}
		matches = append(matches, match[1])
	}
	return matches
}

// ExtractRelativeEndpoints extracts relative endpoints from a data item
func ExtractRelativeEndpoints(data string) []string {
	matches := []string{}
	unique := make(map[string]struct{})

	relativeMatches := relativeEndpointsRegex.FindAllStringSubmatch(data, -1)
	for _, match := range relativeMatches {
		if len(match) < 2 {
			continue
		}
		if _, ok := unique[match[1]]; ok {
			continue
		}

		unique[match[1]] = struct{}{}
		matches = append(matches, match[1])
	}
	return matches
}
