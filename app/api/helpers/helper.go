package helpers

import (
	"net/url"
	"os"
	"strings"
)

// EnforceHTTP ensures the URL has a scheme
func EnforceHTTP(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "http://" + url
}

// RemoveDomainError checks if the URL domain differs from DOMAIN env
func RemoveDomainError(rawURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return true
	} 
	domain := os.Getenv("DOMAIN")
	return parsed.Host != domain
}
