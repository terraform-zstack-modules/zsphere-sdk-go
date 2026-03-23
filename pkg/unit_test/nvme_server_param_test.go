// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateNvmeServerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateNvmeServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateNvmeServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateNvmeServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddNvmeServerParam_MarshalJSON(t *testing.T) {
	p := param.AddNvmeServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddNvmeServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddNvmeServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

