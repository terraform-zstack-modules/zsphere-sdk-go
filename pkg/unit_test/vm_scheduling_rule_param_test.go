// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateVmSchedulingRuleParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVmSchedulingRuleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVmSchedulingRuleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVmSchedulingRuleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateVmSchedulingRuleParam_MarshalJSON(t *testing.T) {
	p := param.CreateVmSchedulingRuleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVmSchedulingRuleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVmSchedulingRuleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

