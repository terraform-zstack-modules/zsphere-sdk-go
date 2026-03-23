// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SnmpAgentInventoryView SnmpAgent
type SnmpAgentInventoryView struct {
	BaseInfoView
	BaseTimeView
	Version string `json:"version,omitempty"`
	ReadCommunity string `json:"readCommunity,omitempty"`
	UserName string `json:"userName,omitempty"`
	AuthAlgorithm string `json:"authAlgorithm,omitempty"`
	AuthPassword string `json:"authPassword,omitempty"`
	PrivacyAlgorithm string `json:"privacyAlgorithm,omitempty"`
	PrivacyPassword string `json:"privacyPassword,omitempty"`
	Port int `json:"port,omitempty"`
	Status string `json:"status,omitempty"`
	SecurityLevel string `json:"securityLevel,omitempty"`
}

// CreateSnmpAgentEventView CreateSnmpAgentEvent
type CreateSnmpAgentEventView struct {
	Inventory SnmpAgentInventoryView `json:"inventory,omitempty"`
}

// StartSnmpAgentEventView StartSnmpAgentEvent
type StartSnmpAgentEventView struct {
	Inventory SnmpAgentInventoryView `json:"inventory,omitempty"`
}

// StopSnmpAgentEventView StopSnmpAgentEvent
type StopSnmpAgentEventView struct {
	Inventory SnmpAgentInventoryView `json:"inventory,omitempty"`
}

// UpdateSnmpAgentEventView UpdateSnmpAgentEvent
type UpdateSnmpAgentEventView struct {
	Inventory SnmpAgentInventoryView `json:"inventory,omitempty"`
}

// QuerySnmpAgentView QuerySnmpAgent
type QuerySnmpAgentView struct {
	Inventories []SnmpAgentInventoryView `json:"inventories,omitempty"`
}

