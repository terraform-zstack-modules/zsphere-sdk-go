// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateExternalPrimaryStorageParam_MarshalJSON(t *testing.T) {
	p := param.UpdateExternalPrimaryStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateExternalPrimaryStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateExternalPrimaryStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddExternalPrimaryStorageParam_MarshalJSON(t *testing.T) {
	p := param.AddExternalPrimaryStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddExternalPrimaryStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddExternalPrimaryStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

