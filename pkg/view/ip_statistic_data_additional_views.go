// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IpStatisticDataView IpStatisticData
type IpStatisticDataView struct {
	Ip string `json:"ip,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	VipName string `json:"vipName,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VmInstanceName string `json:"vmInstanceName,omitempty"`
	VmInstanceType string `json:"vmInstanceType,omitempty"`
	ApplianceVmOwnerUuid string `json:"applianceVmOwnerUuid,omitempty"`
	VmDefaultIp []string `json:"vmDefaultIp,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	State string `json:"state,omitempty"`
	UseFor string `json:"useFor,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	OwnerName string `json:"ownerName,omitempty"`
	ResourceOwnerUuid string `json:"resourceOwnerUuid,omitempty"`
	UsedIpUuid string `json:"usedIpUuid,omitempty"`
}

