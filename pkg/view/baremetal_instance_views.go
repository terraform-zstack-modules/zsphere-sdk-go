// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BaremetalInstanceInventoryView BaremetalInstance
type BaremetalInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	PxeServerUuid string `json:"pxeServerUuid,omitempty"`
	ChassisUuid string `json:"chassisUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	TemplateUuid string `json:"templateUuid,omitempty"`
	Platform string `json:"platform,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Username string `json:"username,omitempty"`
	Port int `json:"port,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	BmNics []BaremetalNicInventoryView `json:"bmNics,omitempty"`
}

// UpdateBaremetalInstanceEventView UpdateBaremetalInstanceEvent
type UpdateBaremetalInstanceEventView struct {
	Inventory BaremetalInstanceInventoryView `json:"inventory,omitempty"`
}

// RebootBaremetalInstanceEventView RebootBaremetalInstanceEvent
type RebootBaremetalInstanceEventView struct {
	Inventory BaremetalInstanceInventoryView `json:"inventory,omitempty"`
}

// StopBaremetalInstanceEventView StopBaremetalInstanceEvent
type StopBaremetalInstanceEventView struct {
	Inventory BaremetalInstanceInventoryView `json:"inventory,omitempty"`
}

// QueryBaremetalInstanceView QueryBaremetalInstance
type QueryBaremetalInstanceView struct {
	Inventories []BaremetalInstanceInventoryView `json:"inventories,omitempty"`
}

// StartBaremetalInstanceEventView StartBaremetalInstanceEvent
type StartBaremetalInstanceEventView struct {
	Inventory BaremetalInstanceInventoryView `json:"inventory,omitempty"`
}

// CreateBaremetalInstanceEventView CreateBaremetalInstanceEvent
type CreateBaremetalInstanceEventView struct {
	Inventory BaremetalInstanceInventoryView `json:"inventory,omitempty"`
}

// DestroyBaremetalInstanceEventView DestroyBaremetalInstanceEvent
type DestroyBaremetalInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// ExpungeBaremetalInstanceEventView ExpungeBaremetalInstanceEvent
type ExpungeBaremetalInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// RecoverBaremetalInstanceEventView RecoverBaremetalInstanceEvent
type RecoverBaremetalInstanceEventView struct {
	Inventory BaremetalInstanceInventoryView `json:"inventory,omitempty"`
}

