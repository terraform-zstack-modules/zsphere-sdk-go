// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateFlkSecSecurityMachineParam_MarshalJSON(t *testing.T) {
	p := param.UpdateFlkSecSecurityMachineParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateFlkSecSecurityMachineParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateFlkSecSecurityMachineParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddFlkSecSecurityMachineParam_MarshalJSON(t *testing.T) {
	p := param.AddFlkSecSecurityMachineParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddFlkSecSecurityMachineParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddFlkSecSecurityMachineParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

