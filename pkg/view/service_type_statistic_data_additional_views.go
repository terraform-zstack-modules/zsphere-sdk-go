// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ServiceTypeStatisticDataView ServiceTypeStatisticData
type ServiceTypeStatisticDataView struct {
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	InterfaceName string `json:"interfaceName,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	ServiceTypes []string `json:"serviceTypes,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	HostName string `json:"hostName,omitempty"`
	HostIp string `json:"hostIp,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

