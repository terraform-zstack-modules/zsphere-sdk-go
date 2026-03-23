// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateDirectoryParam_MarshalJSON(t *testing.T) {
	p := param.UpdateDirectoryParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateDirectoryParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateDirectoryParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateDirectoryParam_MarshalJSON(t *testing.T) {
	p := param.CreateDirectoryParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateDirectoryParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateDirectoryParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

