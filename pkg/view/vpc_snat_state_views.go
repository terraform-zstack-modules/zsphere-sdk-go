// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcSnatStateInventoryView VpcSnatState
type VpcSnatStateInventoryView struct {
	BaseInfoView
	BaseTimeView
	VpcUuid string `json:"vpcUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	State string `json:"state,omitempty"`
}

// QueryVpcSnatStateView QueryVpcSnatState
type QueryVpcSnatStateView struct {
	Inventories []VpcSnatStateInventoryView `json:"inventories,omitempty"`
}

