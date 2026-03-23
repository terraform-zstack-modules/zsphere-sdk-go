// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateVmSchedulingRuleGroupParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVmSchedulingRuleGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVmSchedulingRuleGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVmSchedulingRuleGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateVmSchedulingRuleGroupParam_MarshalJSON(t *testing.T) {
	p := param.CreateVmSchedulingRuleGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVmSchedulingRuleGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVmSchedulingRuleGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

