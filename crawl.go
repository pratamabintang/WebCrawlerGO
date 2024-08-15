package WebCrawlerGO

import (
	"net"
	"net/url"
	"strings"
)

type customError struct {
	funcError, msg string
}

func (e *customError) Error() string {
	return e.funcError + ": " + e.msg
}

func NormalizeUrl(inputURL string) (string, error) {
	if rawURL, err := url.ParseRequestURI(inputURL); err != nil {
		return "", &customError{"NormalizeURL error", "URL not valid"}
	} else {
		host, _, noPort := net.SplitHostPort(rawURL.Host)
		if noPort != nil && !strings.Contains(rawURL.Host, ":") {
			host = rawURL.Host
		}

		validURL := rawURL.Scheme + "://" + host + rawURL.Path
		if len(validURL) > 0 && validURL[len(validURL)-1] == '/' {
			return validURL[:len(validURL)-1], nil
		}

		return validURL, nil
	}
}
