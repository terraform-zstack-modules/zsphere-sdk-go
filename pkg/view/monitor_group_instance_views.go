// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorGroupInstanceInventoryView MonitorGroupInstance
type MonitorGroupInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	GroupUuid string `json:"groupUuid,omitempty"`
	InstanceResourceType string `json:"instanceResourceType,omitempty"`
	InstanceUuid string `json:"instanceUuid,omitempty"`
	Status string `json:"status,omitempty"`
}

// AddInstanceToMonitorGroupEventView AddInstanceToMonitorGroupEvent
type AddInstanceToMonitorGroupEventView struct {
	Inventory MonitorGroupInstanceInventoryView `json:"inventory,omitempty"`
}

// QueryMonitorGroupInstanceView QueryMonitorGroupInstance
type QueryMonitorGroupInstanceView struct {
	Inventories []MonitorGroupInstanceInventoryView `json:"inventories,omitempty"`
}

