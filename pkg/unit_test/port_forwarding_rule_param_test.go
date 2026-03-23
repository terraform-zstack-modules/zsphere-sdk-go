// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdatePortForwardingRuleParam_MarshalJSON(t *testing.T) {
	p := param.UpdatePortForwardingRuleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdatePortForwardingRuleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdatePortForwardingRuleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreatePortForwardingRuleParam_MarshalJSON(t *testing.T) {
	p := param.CreatePortForwardingRuleParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreatePortForwardingRuleParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreatePortForwardingRuleParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

