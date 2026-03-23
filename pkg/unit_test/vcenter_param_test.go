// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestAddVCenterParam_MarshalJSON(t *testing.T) {
	p := param.AddVCenterParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddVCenterParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddVCenterParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateVCenterParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVCenterParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVCenterParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVCenterParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

