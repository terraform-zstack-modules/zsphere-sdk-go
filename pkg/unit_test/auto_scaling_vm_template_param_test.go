// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateAutoScalingVmTemplateParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAutoScalingVmTemplateParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAutoScalingVmTemplateParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAutoScalingVmTemplateParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateAutoScalingVmTemplateParam_MarshalJSON(t *testing.T) {
	p := param.CreateAutoScalingVmTemplateParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateAutoScalingVmTemplateParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateAutoScalingVmTemplateParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

