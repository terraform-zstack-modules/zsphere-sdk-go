// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateBlockPrimaryStorageParam_MarshalJSON(t *testing.T) {
	p := param.UpdateBlockPrimaryStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateBlockPrimaryStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateBlockPrimaryStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddBlockPrimaryStorageParam_MarshalJSON(t *testing.T) {
	p := param.AddBlockPrimaryStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddBlockPrimaryStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddBlockPrimaryStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

