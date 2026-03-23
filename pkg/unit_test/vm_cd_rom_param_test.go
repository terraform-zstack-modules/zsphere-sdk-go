// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateVmCdRomParam_MarshalJSON(t *testing.T) {
	p := param.CreateVmCdRomParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVmCdRomParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVmCdRomParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateVmCdRomParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVmCdRomParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVmCdRomParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVmCdRomParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

