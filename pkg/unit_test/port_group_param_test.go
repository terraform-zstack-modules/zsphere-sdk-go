// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreatePortGroupParam_MarshalJSON(t *testing.T) {
	p := param.CreatePortGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreatePortGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreatePortGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdatePortGroupParam_MarshalJSON(t *testing.T) {
	p := param.UpdatePortGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdatePortGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdatePortGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

