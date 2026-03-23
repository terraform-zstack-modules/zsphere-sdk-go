// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateCdpPolicyParam_MarshalJSON(t *testing.T) {
	p := param.CreateCdpPolicyParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateCdpPolicyParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateCdpPolicyParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateCdpPolicyParam_MarshalJSON(t *testing.T) {
	p := param.UpdateCdpPolicyParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateCdpPolicyParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateCdpPolicyParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

