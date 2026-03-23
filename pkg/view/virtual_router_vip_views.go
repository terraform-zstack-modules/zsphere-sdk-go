// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterVipInventoryView VirtualRouterVip
type VirtualRouterVipInventoryView struct {
	BaseInfoView
	BaseTimeView
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid,omitempty"`
}

