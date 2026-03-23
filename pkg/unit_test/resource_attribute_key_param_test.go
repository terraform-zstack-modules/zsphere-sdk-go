// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateResourceAttributeKeyParam_MarshalJSON(t *testing.T) {
	p := param.UpdateResourceAttributeKeyParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateResourceAttributeKeyParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateResourceAttributeKeyParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateResourceAttributeKeyParam_MarshalJSON(t *testing.T) {
	p := param.CreateResourceAttributeKeyParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateResourceAttributeKeyParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateResourceAttributeKeyParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

