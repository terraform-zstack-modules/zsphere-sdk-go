// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateVolumeBackupParam_MarshalJSON(t *testing.T) {
	p := param.CreateVolumeBackupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVolumeBackupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVolumeBackupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

