// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateBaremetalPxeServerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateBaremetalPxeServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateBaremetalPxeServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateBaremetalPxeServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateBaremetalPxeServerParam_MarshalJSON(t *testing.T) {
	p := param.CreateBaremetalPxeServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateBaremetalPxeServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateBaremetalPxeServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

