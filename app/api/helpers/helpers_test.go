package helpers_test

import (
	"os"
	"testing"

	"github.com/chethanm99/go-url-shortner/api/helpers"
)

func TestEnforceHTTP(t *testing.T) {
	got := helpers.EnforceHTTP("example.com")
	want := "http://example.com"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRemoveDomainError(t *testing.T) {
	os.Setenv("DOMAIN", "example.com")

	if helpers.RemoveDomainError("http://example.com") {
		t.Errorf("expected false when domain matches")
	}

	if !helpers.RemoveDomainError("http://google.com") {
		t.Errorf("expected true when domain is different")
	}
}
