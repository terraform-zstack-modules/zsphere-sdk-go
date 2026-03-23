// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateVpcHaGroupParam_MarshalJSON(t *testing.T) {
	p := param.CreateVpcHaGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVpcHaGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVpcHaGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateVpcHaGroupParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVpcHaGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVpcHaGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVpcHaGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

