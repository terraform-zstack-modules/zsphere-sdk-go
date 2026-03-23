// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ConsoleProxyAgentInventoryView ConsoleProxyAgent
type ConsoleProxyAgentInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	ConsoleProxyOverriddenIp string `json:"consoleProxyOverriddenIp,omitempty"`
	ConsoleProxyPort int `json:"consoleProxyPort,omitempty"`
	Type string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
}

// ReconnectConsoleProxyAgentEventView ReconnectConsoleProxyAgentEvent
type ReconnectConsoleProxyAgentEventView struct {
	Inventory map[string]interface{} `json:"inventory,omitempty"`
}

// QueryConsoleProxyAgentView QueryConsoleProxyAgent
type QueryConsoleProxyAgentView struct {
	Inventories []ConsoleProxyAgentInventoryView `json:"inventories,omitempty"`
}

// UpdateConsoleProxyAgentEventView UpdateConsoleProxyAgentEvent
type UpdateConsoleProxyAgentEventView struct {
	Inventory ConsoleProxyAgentInventoryView `json:"inventory,omitempty"`
}

