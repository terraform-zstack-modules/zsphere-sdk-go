// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestAddIscsiServerParam_MarshalJSON(t *testing.T) {
	p := param.AddIscsiServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddIscsiServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddIscsiServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateIscsiServerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateIscsiServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateIscsiServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateIscsiServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

