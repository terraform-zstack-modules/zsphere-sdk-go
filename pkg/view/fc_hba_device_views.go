// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FcHbaDeviceInventoryView FcHbaDevice
type FcHbaDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	PortName string `json:"portName,omitempty"`
	PortState string `json:"portState,omitempty"`
	Speed string `json:"speed,omitempty"`
	SupportedSpeeds string `json:"supportedSpeeds,omitempty"`
	SymbolicName string `json:"symbolicName,omitempty"`
	SupportedClasses string `json:"supportedClasses,omitempty"`
	NodeName string `json:"nodeName,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	HbaType string `json:"hbaType,omitempty"`
}

// QueryFcHbaDeviceView QueryFcHbaDevice
type QueryFcHbaDeviceView struct {
	Inventories []HbaDeviceInventoryView `json:"inventories,omitempty"`
}

