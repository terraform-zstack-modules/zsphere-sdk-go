// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestAddEventRuleTemplateParam_MarshalJSON(t *testing.T) {
	p := param.AddEventRuleTemplateParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddEventRuleTemplateParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddEventRuleTemplateParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateEventRuleTemplateParam_MarshalJSON(t *testing.T) {
	p := param.UpdateEventRuleTemplateParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateEventRuleTemplateParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateEventRuleTemplateParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

