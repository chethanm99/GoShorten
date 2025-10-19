package database

import (
	"testing"
)

func TestConnect(t *testing.T) {
	t.Log("Testing database connection using Connect() function...")

	err := Connect()

	if err != nil {
		t.Fatalf("FAIL: Database connection failed. Error: %v", err)
	}

	t.Log("PASS: Database connection successful (Connect() returned nil).")
}
