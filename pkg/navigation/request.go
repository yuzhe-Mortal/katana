package navigation

import (
	"strings"
)

// Depth is the depth of a navigation
type Depth struct{}

// Request is a navigation request for the crawler
type Request struct {
	Method       string
	URL          string
	Body         string
	Depth        int
	Headers      map[string]string
	Tag          string
	Attribute    string
	RootHostname string
	Source       string // source is the source of the request
}

// RequestURL returns the request URL for the navigation
func (n *Request) RequestURL() string {
	switch n.Method {
	case "GET":
		return n.URL
	case "POST":
		builder := &strings.Builder{}
		builder.WriteString(n.URL)
		builder.WriteString(":")
		builder.WriteString(n.Body)
		builtURL := builder.String()
		return builtURL
	}
	return ""
}

// newNavigationRequestURL generates a navigation request from a relative URL
func NewNavigationRequestURLFromResponse(path, source, tag, attribute string, resp Response) Request {
	requestURL := resp.AbsoluteURL(path) //通过相对路径生成绝对路径
	return Request{Method: "GET", URL: requestURL, RootHostname: resp.RootHostname, Depth: resp.Depth, Source: source, Attribute: attribute, Tag: tag}
}
