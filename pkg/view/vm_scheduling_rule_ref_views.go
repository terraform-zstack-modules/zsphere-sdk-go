// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmSchedulingRuleRefInventoryView VmSchedulingRuleRef
type VmSchedulingRuleRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmGroupUuid string `json:"vmGroupUuid,omitempty"`
	HostGroupUuid string `json:"hostGroupUuid,omitempty"`
	VmSchedulingRuleUuid string `json:"vmSchedulingRuleUuid,omitempty"`
}

