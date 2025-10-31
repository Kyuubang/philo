package api

import (
	"testing"
)

// TestHTTPClientReuse verifies that getHTTPClient returns the same instance
func TestHTTPClientReuse(t *testing.T) {
	// Get the first client instance
	client1 := getHTTPClient()
	if client1 == nil {
		t.Fatal("getHTTPClient returned nil")
	}

	// Get a second client instance
	client2 := getHTTPClient()
	if client2 == nil {
		t.Fatal("getHTTPClient returned nil on second call")
	}

	// Verify they are the same instance (same memory address)
	if client1 != client2 {
		t.Error("getHTTPClient should return the same HTTP client instance for connection reuse")
	}

	// Get a third client instance to be thorough
	client3 := getHTTPClient()
	if client1 != client3 {
		t.Error("getHTTPClient should consistently return the same HTTP client instance")
	}

	t.Log("HTTP client reuse verified successfully")
}
