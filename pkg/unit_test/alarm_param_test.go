// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateAlarmParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAlarmParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAlarmParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAlarmParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateAlarmParam_MarshalJSON(t *testing.T) {
	p := param.CreateAlarmParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateAlarmParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateAlarmParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

