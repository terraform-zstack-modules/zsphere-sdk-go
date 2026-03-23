// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateVniRangeParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVniRangeParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVniRangeParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVniRangeParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateVniRangeParam_MarshalJSON(t *testing.T) {
	p := param.CreateVniRangeParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVniRangeParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVniRangeParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

