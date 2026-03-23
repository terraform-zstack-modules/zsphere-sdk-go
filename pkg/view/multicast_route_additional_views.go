// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MulticastRouteInventoryView MulticastRoute
type MulticastRouteInventoryView struct {
	BaseInfoView
	BaseTimeView
	SourceAddress string `json:"sourceAddress,omitempty"`
	GroupAddress string `json:"groupAddress,omitempty"`
	IngressInterfaces string `json:"ingressInterfaces,omitempty"`
	EgressInterfaces string `json:"egressInterfaces,omitempty"`
}

