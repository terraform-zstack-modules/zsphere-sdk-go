// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestCreateVolumeSnapshotGroupParam_MarshalJSON(t *testing.T) {
	p := param.CreateVolumeSnapshotGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateVolumeSnapshotGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateVolumeSnapshotGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateVolumeSnapshotGroupParam_MarshalJSON(t *testing.T) {
	p := param.UpdateVolumeSnapshotGroupParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateVolumeSnapshotGroupParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateVolumeSnapshotGroupParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

