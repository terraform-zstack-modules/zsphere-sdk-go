// Copyright (c) ZStack.io, Inc.
// Auto-generated test infrastructure. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/client"
)

// newMockClient creates a ZSClient backed by an httptest server.
// The handler receives all HTTP requests and can assert on method/path/body.
func newMockClient(handler http.HandlerFunc) (*client.ZSClient, func()) {
	server := httptest.NewServer(handler)
	// Parse host and port from server URL
	addr := server.Listener.Addr().String()
	parts := strings.SplitN(addr, ":", 2)
	host := parts[0]
	port := 80
	if len(parts) == 2 {
		fmt.Sscanf(parts[1], "%d", &port)
	}

	config := client.NewZSConfig(host, port, "")
	config.LoginAccount("admin", "password")
	cli := client.NewZSClient(config)
	cli.LoadSession("mock-session-id")
	return cli, server.Close
}

// mockInventoryResponse builds a JSON response wrapping data in {"inventory": ...}
func mockInventoryResponse(data map[string]interface{}) []byte {
	resp := map[string]interface{}{"inventory": data}
	b, _ := json.Marshal(resp)
	return b
}

// mockInventoriesResponse builds a JSON response wrapping data in {"inventories": [...]}
func mockInventoriesResponse(items ...map[string]interface{}) []byte {
	resp := map[string]interface{}{"inventories": items}
	b, _ := json.Marshal(resp)
	return b
}

// stringPtr returns a pointer to the given string.
func stringPtr(s string) *string {
	return &s
}

// timePtr parses a time string and returns a pointer.
func timePtr(s string) *time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return &t
}

// assertEqual is a simple test helper.
func assertEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

// assertNoError fails the test if err is not nil.
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// assertContains checks that s contains substr.
func assertContains(t *testing.T, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("expected %q to contain %q", s, substr)
	}
}
