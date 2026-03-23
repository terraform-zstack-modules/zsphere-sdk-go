// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestAddKVMHostParam_MarshalJSON(t *testing.T) {
	p := param.AddKVMHostParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddKVMHostParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddKVMHostParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateKVMHostParam_MarshalJSON(t *testing.T) {
	p := param.UpdateKVMHostParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateKVMHostParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateKVMHostParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

