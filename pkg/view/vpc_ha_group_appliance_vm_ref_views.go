// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcHaGroupApplianceVmRefInventoryView VpcHaGroupApplianceVmRef
type VpcHaGroupApplianceVmRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VpcHaRouterUuid string `json:"vpcHaRouterUuid,omitempty"`
}

