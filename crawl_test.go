package WebCrawlerGO

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	test := []struct {
		name, inputURL, expected string
	}{
		{"1", "https://chatgpt.com/?oai-dm=1", "https://chatgpt.com"},
		{"2", "https://wagslane.dev", "https://wagslane.dev"},
		{"3", "https://api.example.com:3000/users/12345?status=active&sort=desc#details", "https://api.example.com/users/12345"},
		{"4", "https://wagslane.dev/posts/a-case-against-a-case-for-the-book-of-mormon/", "https://wagslane.dev/posts/a-case-against-a-case-for-the-book-of-mormon"},
		{"5", "https://wagslane.dev/posts/guard-keyword-error-handling-golang/", "https://wagslane.dev/posts/guard-keyword-error-handling-golang"},
		{"6", "https://wagslane.dev/posts/college-a-solution-in-search-of-a-problem/", "https://wagslane.dev/posts/college-a-solution-in-search-of-a-problem"},
		{"7", "https://wagslane.dev/posts/continuous-deployments-arent-continuous-disruptions", "https://wagslane.dev/posts/continuous-deployments-arent-continuous-disruptions"},
	}

	for _, testcase := range test {
		t.Run(testcase.name, func(t *testing.T) {
			result, err := NormalizeUrl(testcase.inputURL)
			assert.NoError(t, err, "must be valid and pass error test")
			assert.Equalf(t, testcase.expected, result, "wrong result")
		})
	}

}
