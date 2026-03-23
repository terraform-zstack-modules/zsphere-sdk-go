// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MulticastRouterVpcVRouterRefInventoryView MulticastRouterVpcVRouterRef
type MulticastRouterVpcVRouterRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VpcRouterUuid string `json:"vpcRouterUuid,omitempty"`
}

