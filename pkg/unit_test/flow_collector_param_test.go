// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateFlowCollectorParam_MarshalJSON(t *testing.T) {
	p := param.CreateFlowCollectorParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateFlowCollectorParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateFlowCollectorParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateFlowCollectorParam_MarshalJSON(t *testing.T) {
	p := param.UpdateFlowCollectorParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateFlowCollectorParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateFlowCollectorParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

