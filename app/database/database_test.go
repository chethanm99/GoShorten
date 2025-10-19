package database

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	RDB = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	m.Run()
}
func TestConnect(t *testing.T) {
	t.Log("Testing database connection using Connect() function...")

	err := Connect()

	if err != nil {
		t.Fatalf("FAIL: Database connection failed. Error: %v", err)
	}

	t.Log("PASS: Database connection successful (Connect() returned nil).")
}
