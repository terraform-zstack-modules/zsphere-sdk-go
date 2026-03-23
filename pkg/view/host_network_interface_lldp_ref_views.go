// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostNetworkInterfaceLldpRefInventoryView HostNetworkInterfaceLldpRef
type HostNetworkInterfaceLldpRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	LldpUuid string `json:"lldpUuid,omitempty"`
	ChassisId string `json:"chassisId,omitempty"`
	TimeToLive int `json:"timeToLive,omitempty"`
	ManagementAddress string `json:"managementAddress,omitempty"`
	SystemName string `json:"systemName,omitempty"`
	SystemDescription string `json:"systemDescription,omitempty"`
	SystemCapabilities string `json:"systemCapabilities,omitempty"`
	PortId string `json:"portId,omitempty"`
	PortDescription string `json:"portDescription,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	AggregationPortId int64 `json:"aggregationPortId,omitempty"`
	Mtu int `json:"mtu,omitempty"`
}

