// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateFlowMeterParam_MarshalJSON(t *testing.T) {
	p := param.CreateFlowMeterParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateFlowMeterParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateFlowMeterParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateFlowMeterParam_MarshalJSON(t *testing.T) {
	p := param.UpdateFlowMeterParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateFlowMeterParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateFlowMeterParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

