// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateSchedulerJobParam_MarshalJSON(t *testing.T) {
	p := param.CreateSchedulerJobParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSchedulerJobParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSchedulerJobParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateSchedulerJobParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSchedulerJobParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSchedulerJobParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSchedulerJobParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

