// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateLoadBalancerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateLoadBalancerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateLoadBalancerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateLoadBalancerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateLoadBalancerParam_MarshalJSON(t *testing.T) {
	p := param.CreateLoadBalancerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateLoadBalancerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateLoadBalancerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

