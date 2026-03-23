// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterEipRefInventoryView VirtualRouterEipRef
type VirtualRouterEipRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	EipUuid string `json:"eipUuid,omitempty"`
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid,omitempty"`
}

