// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ManagementNodeInventoryView ManagementNode
type ManagementNodeInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostName string `json:"hostName,omitempty"`
	JoinDate time.Time `json:"joinDate,omitempty"`
	HeartBeat time.Time `json:"heartBeat,omitempty"`
}

// QueryManagementNodeView QueryManagementNode
type QueryManagementNodeView struct {
	Inventories []ManagementNodeInventoryView `json:"inventories,omitempty"`
}

