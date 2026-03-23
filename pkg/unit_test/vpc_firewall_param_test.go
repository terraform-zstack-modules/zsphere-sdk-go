// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateVpcFirewallParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVpcFirewallParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVpcFirewallParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVpcFirewallParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateVpcFirewallParam_MarshalJSON(t *testing.T) {
	p := param.CreateVpcFirewallParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVpcFirewallParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVpcFirewallParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

