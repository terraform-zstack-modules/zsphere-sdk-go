// Copyright (c) ZStack.io, Inc.
// Auto-generated view tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

func TestVCenterPrimaryStorageInventoryView_UnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"uuid": "test-uuid-001",
		"name": "test-vcenter_primary_storage",
		"createDate": "2024-01-01T00:00:00.000+08:00",
		"lastOpDate": "2024-01-01T00:00:00.000+08:00"
	}`
	var v view.VCenterPrimaryStorageInventoryView
	err := json.Unmarshal([]byte(jsonStr), &v)
	assertNoError(t, err)
}

func TestVCenterPrimaryStorageInventoryView_UnmarshalEmpty(t *testing.T) {
	var v view.VCenterPrimaryStorageInventoryView
	err := json.Unmarshal([]byte(`{}`), &v)
	assertNoError(t, err)
}

