package WebCrawlerGO

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	result, err := NormalizeUrl("https://chatgpt.com/?oai-dm=1")
	assert.NoError(t, err, "must be valid and pass error test")
	assert.Equalf(t, "https://chatgpt.com", result, "wrong result")

	result, err = NormalizeUrl("https://wagslane.dev")
	assert.NoError(t, err, "must be valid and pass error test")
	assert.Equalf(t, "https://wagslane.dev", result, "wrong result")

	result, err = NormalizeUrl("https://api.example.com:3000/users/12345?status=active&sort=desc#details")
	assert.NoError(t, err, "must be valid and pass error test")
	assert.Equalf(t, "https://api.example.com/users/12345", result, "wrong result")

	result, err = NormalizeUrl("https://wagslane.dev/posts/a-case-against-a-case-for-the-book-of-mormon/")
	assert.NoError(t, err, "must be valid and pass error test")
	assert.Equalf(t, "https://wagslane.dev/posts/a-case-against-a-case-for-the-book-of-mormon", result, "wrong result")

	result, err = NormalizeUrl("https://wagslane.dev/posts/guard-keyword-error-handling-golang/")
	assert.NoError(t, err, "must be valid and pass error test")
	assert.Equalf(t, "https://wagslane.dev/posts/guard-keyword-error-handling-golang", result, "wrong result")

	result, err = NormalizeUrl("https://wagslane.dev/posts/college-a-solution-in-search-of-a-problem/")
	assert.NoError(t, err, "must be valid and pass error test")
	assert.Equalf(t, "https://wagslane.dev/posts/college-a-solution-in-search-of-a-problem", result, "wrong result")

	result, err = NormalizeUrl("https://wagslane.dev/posts/continuous-deployments-arent-continuous-disruptions")
	assert.NoError(t, err, "must be valid and pass error test")
	assert.Equalf(t, "https://wagslane.dev/posts/continuous-deployments-arent-continuous-disruptions", result, "wrong result")
}
