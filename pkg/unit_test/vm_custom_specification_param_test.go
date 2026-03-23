// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateVmCustomSpecificationParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVmCustomSpecificationParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVmCustomSpecificationParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVmCustomSpecificationParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateVmCustomSpecificationParam_MarshalJSON(t *testing.T) {
	p := param.CreateVmCustomSpecificationParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVmCustomSpecificationParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVmCustomSpecificationParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

