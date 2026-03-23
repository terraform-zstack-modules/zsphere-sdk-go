// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateZoneParam_MarshalJSON(t *testing.T) {
	p := param.UpdateZoneParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateZoneParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateZoneParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateZoneParam_MarshalJSON(t *testing.T) {
	p := param.CreateZoneParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateZoneParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateZoneParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

