// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateSchedulerTriggerParam_MarshalJSON(t *testing.T) {
	p := param.CreateSchedulerTriggerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSchedulerTriggerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSchedulerTriggerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateSchedulerTriggerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSchedulerTriggerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSchedulerTriggerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSchedulerTriggerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

