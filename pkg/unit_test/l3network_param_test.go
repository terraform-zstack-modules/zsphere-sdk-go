// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateL3NetworkParam_MarshalJSON(t *testing.T) {
	p := param.CreateL3NetworkParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateL3NetworkParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateL3NetworkParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateL3NetworkParam_MarshalJSON(t *testing.T) {
	p := param.UpdateL3NetworkParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateL3NetworkParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateL3NetworkParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

