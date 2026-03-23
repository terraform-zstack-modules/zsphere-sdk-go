// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateLoadBalancerListenerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateLoadBalancerListenerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateLoadBalancerListenerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateLoadBalancerListenerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateLoadBalancerListenerParam_MarshalJSON(t *testing.T) {
	p := param.CreateLoadBalancerListenerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateLoadBalancerListenerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateLoadBalancerListenerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

