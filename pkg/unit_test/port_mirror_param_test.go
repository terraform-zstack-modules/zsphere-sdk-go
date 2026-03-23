// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreatePortMirrorParam_MarshalJSON(t *testing.T) {
	p := param.CreatePortMirrorParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreatePortMirrorParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreatePortMirrorParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdatePortMirrorParam_MarshalJSON(t *testing.T) {
	p := param.UpdatePortMirrorParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdatePortMirrorParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdatePortMirrorParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

