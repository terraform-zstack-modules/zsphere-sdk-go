// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BaremetalPxeServerInventoryView BaremetalPxeServer
type BaremetalPxeServerInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Description string `json:"description,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	SshUsername string `json:"sshUsername,omitempty"`
	SshPassword string `json:"sshPassword,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	StoragePath string `json:"storagePath,omitempty"`
	DhcpInterface string `json:"dhcpInterface,omitempty"`
	DhcpInterfaceAddress string `json:"dhcpInterfaceAddress,omitempty"`
	DhcpRangeBegin string `json:"dhcpRangeBegin,omitempty"`
	DhcpRangeEnd string `json:"dhcpRangeEnd,omitempty"`
	DhcpRangeNetmask string `json:"dhcpRangeNetmask,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// ReconnectBaremetalPxeServerEventView ReconnectBaremetalPxeServerEvent
type ReconnectBaremetalPxeServerEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

// DeleteBaremetalPxeServerEventView DeleteBaremetalPxeServerEvent
type DeleteBaremetalPxeServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachBaremetalPxeServerToClusterEventView AttachBaremetalPxeServerToClusterEvent
type AttachBaremetalPxeServerToClusterEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

// StopBaremetalPxeServerEventView StopBaremetalPxeServerEvent
type StopBaremetalPxeServerEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

// UpdateBaremetalPxeServerEventView UpdateBaremetalPxeServerEvent
type UpdateBaremetalPxeServerEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

// StartBaremetalPxeServerEventView StartBaremetalPxeServerEvent
type StartBaremetalPxeServerEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

// QueryBaremetalPxeServerView QueryBaremetalPxeServer
type QueryBaremetalPxeServerView struct {
	Inventories []BaremetalPxeServerInventoryView `json:"inventories,omitempty"`
}

// DetachBaremetalPxeServerFromClusterEventView DetachBaremetalPxeServerFromClusterEvent
type DetachBaremetalPxeServerFromClusterEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

// CreateBaremetalPxeServerEventView CreateBaremetalPxeServerEvent
type CreateBaremetalPxeServerEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

