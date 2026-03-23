// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateWebhookParam_MarshalJSON(t *testing.T) {
	p := param.CreateWebhookParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateWebhookParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateWebhookParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateWebhookParam_MarshalJSON(t *testing.T) {
	p := param.UpdateWebhookParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateWebhookParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateWebhookParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

