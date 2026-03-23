// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ReservedIpRangeInventoryView ReservedIpRange
type ReservedIpRangeInventoryView struct {
	BaseInfoView
	BaseTimeView
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Description string `json:"description,omitempty"`
	StartIp string `json:"startIp,omitempty"`
	EndIp string `json:"endIp,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
}

// DeleteReservedIpRangeEventView DeleteReservedIpRangeEvent
type DeleteReservedIpRangeEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddReservedIpRangeEventView AddReservedIpRangeEvent
type AddReservedIpRangeEventView struct {
	Inventory ReservedIpRangeInventoryView `json:"inventory,omitempty"`
}

