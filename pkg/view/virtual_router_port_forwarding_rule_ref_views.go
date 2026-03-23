// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterPortForwardingRuleRefInventoryView VirtualRouterPortForwardingRuleRef
type VirtualRouterPortForwardingRuleRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VipUuid string `json:"vipUuid,omitempty"`
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid,omitempty"`
}

