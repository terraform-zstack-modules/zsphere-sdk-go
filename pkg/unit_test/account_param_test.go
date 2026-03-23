// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateAccountParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAccountParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAccountParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAccountParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateAccountParam_MarshalJSON(t *testing.T) {
	p := param.CreateAccountParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateAccountParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateAccountParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

