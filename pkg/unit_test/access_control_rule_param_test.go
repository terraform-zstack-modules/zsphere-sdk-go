// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestAddAccessControlRuleParam_MarshalJSON(t *testing.T) {
	p := param.AddAccessControlRuleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddAccessControlRuleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddAccessControlRuleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateAccessControlRuleParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAccessControlRuleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAccessControlRuleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAccessControlRuleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

