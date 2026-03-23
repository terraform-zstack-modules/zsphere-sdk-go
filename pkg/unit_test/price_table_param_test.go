// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreatePriceTableParam_MarshalJSON(t *testing.T) {
	p := param.CreatePriceTableParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreatePriceTableParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreatePriceTableParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdatePriceTableParam_MarshalJSON(t *testing.T) {
	p := param.UpdatePriceTableParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdatePriceTableParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdatePriceTableParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

