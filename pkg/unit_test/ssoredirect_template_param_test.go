// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateSSORedirectTemplateParam_MarshalJSON(t *testing.T) {
	p := param.CreateSSORedirectTemplateParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSSORedirectTemplateParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSSORedirectTemplateParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateSSORedirectTemplateParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSSORedirectTemplateParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSSORedirectTemplateParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSSORedirectTemplateParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

