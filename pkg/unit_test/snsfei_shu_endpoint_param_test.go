// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateSNSFeiShuEndpointParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSNSFeiShuEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSNSFeiShuEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSNSFeiShuEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateSNSFeiShuEndpointParam_MarshalJSON(t *testing.T) {
	p := param.CreateSNSFeiShuEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSNSFeiShuEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSNSFeiShuEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

