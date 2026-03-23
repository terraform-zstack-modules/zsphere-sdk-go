// Copyright (c) ZStack.io, Inc.
// Auto-generated view tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

func TestSNSHttpEndpointInventoryView_UnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"uuid": "test-uuid-001",
		"name": "test-snshttp_endpoint",
		"createDate": "2024-01-01T00:00:00.000+08:00",
		"lastOpDate": "2024-01-01T00:00:00.000+08:00"
	}`
	var v view.SNSHttpEndpointInventoryView
	err := json.Unmarshal([]byte(jsonStr), &v)
	assertNoError(t, err)
}

func TestSNSHttpEndpointInventoryView_UnmarshalEmpty(t *testing.T) {
	var v view.SNSHttpEndpointInventoryView
	err := json.Unmarshal([]byte(`{}`), &v)
	assertNoError(t, err)
}

