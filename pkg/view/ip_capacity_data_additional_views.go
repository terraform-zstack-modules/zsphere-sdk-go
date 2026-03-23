// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IpCapacityDataView IpCapacityData
type IpCapacityDataView struct {
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	UsedIpAddressNumber int64 `json:"usedIpAddressNumber,omitempty"`
	Ipv4TotalCapacity int64 `json:"ipv4TotalCapacity,omitempty"`
	Ipv4AvailableCapacity int64 `json:"ipv4AvailableCapacity,omitempty"`
	Ipv4UsedIpAddressNumber int64 `json:"ipv4UsedIpAddressNumber,omitempty"`
	Ipv6TotalCapacity int64 `json:"ipv6TotalCapacity,omitempty"`
	Ipv6AvailableCapacity int64 `json:"ipv6AvailableCapacity,omitempty"`
	Ipv6UsedIpAddressNumber int64 `json:"ipv6UsedIpAddressNumber,omitempty"`
}

