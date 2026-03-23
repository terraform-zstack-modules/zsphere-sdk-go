// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateSNSHttpEndpointParam_MarshalJSON(t *testing.T) {
	p := param.CreateSNSHttpEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSNSHttpEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSNSHttpEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateSNSHttpEndpointParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSNSHttpEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSNSHttpEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSNSHttpEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

