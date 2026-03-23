// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateThirdpartyPlatformParam_MarshalJSON(t *testing.T) {
	p := param.UpdateThirdpartyPlatformParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateThirdpartyPlatformParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateThirdpartyPlatformParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddThirdpartyPlatformParam_MarshalJSON(t *testing.T) {
	p := param.AddThirdpartyPlatformParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddThirdpartyPlatformParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddThirdpartyPlatformParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

