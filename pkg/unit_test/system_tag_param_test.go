// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateSystemTagParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSystemTagParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSystemTagParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSystemTagParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateSystemTagParam_MarshalJSON(t *testing.T) {
	p := param.CreateSystemTagParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSystemTagParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSystemTagParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

