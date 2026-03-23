// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateRoleParam_MarshalJSON(t *testing.T) {
	p := param.CreateRoleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateRoleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateRoleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateRoleParam_MarshalJSON(t *testing.T) {
	p := param.UpdateRoleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateRoleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateRoleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

