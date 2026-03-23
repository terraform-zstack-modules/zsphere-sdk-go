// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateSNSSnmpPlatformParam_MarshalJSON(t *testing.T) {
	p := param.CreateSNSSnmpPlatformParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSNSSnmpPlatformParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSNSSnmpPlatformParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateSNSSnmpPlatformParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSNSSnmpPlatformParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSNSSnmpPlatformParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSNSSnmpPlatformParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

