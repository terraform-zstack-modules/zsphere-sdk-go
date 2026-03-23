// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestAddIpRangeParam_MarshalJSON(t *testing.T) {
	p := param.AddIpRangeParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddIpRangeParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddIpRangeParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateIpRangeParam_MarshalJSON(t *testing.T) {
	p := param.UpdateIpRangeParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateIpRangeParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateIpRangeParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

