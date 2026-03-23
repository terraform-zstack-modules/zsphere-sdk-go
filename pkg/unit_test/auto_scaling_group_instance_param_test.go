// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestUpdateAutoScalingGroupInstanceParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAutoScalingGroupInstanceParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAutoScalingGroupInstanceParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAutoScalingGroupInstanceParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

