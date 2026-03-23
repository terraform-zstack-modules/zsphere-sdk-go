// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateBaremetalInstanceParam_MarshalJSON(t *testing.T) {
	p := param.UpdateBaremetalInstanceParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateBaremetalInstanceParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateBaremetalInstanceParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateBaremetalInstanceParam_MarshalJSON(t *testing.T) {
	p := param.CreateBaremetalInstanceParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateBaremetalInstanceParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateBaremetalInstanceParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

