// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestAddLdapServerParam_MarshalJSON(t *testing.T) {
	p := param.AddLdapServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddLdapServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddLdapServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateLdapServerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateLdapServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateLdapServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateLdapServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

