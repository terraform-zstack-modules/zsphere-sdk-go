// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MulticastRouterRendezvousPointInventoryView MulticastRouterRendezvousPoint
type MulticastRouterRendezvousPointInventoryView struct {
	BaseInfoView
	BaseTimeView
	MulticastRouterUuid string `json:"multicastRouterUuid,omitempty"`
	RpAddress string `json:"rpAddress,omitempty"`
	GroupAddress string `json:"groupAddress,omitempty"`
}

