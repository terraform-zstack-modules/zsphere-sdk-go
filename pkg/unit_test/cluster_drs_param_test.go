// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateClusterDRSParam_MarshalJSON(t *testing.T) {
	p := param.CreateClusterDRSParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateClusterDRSParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateClusterDRSParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateClusterDRSParam_MarshalJSON(t *testing.T) {
	p := param.UpdateClusterDRSParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateClusterDRSParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateClusterDRSParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

