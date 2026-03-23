// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostKernelInterfaceTrafficTypeInventoryView HostKernelInterfaceTrafficType
type HostKernelInterfaceTrafficTypeInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostKernelInterfaceUuid string `json:"hostKernelInterfaceUuid,omitempty"`
	TrafficType string `json:"trafficType,omitempty"`
}

