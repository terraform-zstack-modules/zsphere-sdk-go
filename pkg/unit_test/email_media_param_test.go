// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateEmailMediaParam_MarshalJSON(t *testing.T) {
	p := param.CreateEmailMediaParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateEmailMediaParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateEmailMediaParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateEmailMediaParam_MarshalJSON(t *testing.T) {
	p := param.UpdateEmailMediaParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateEmailMediaParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateEmailMediaParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

