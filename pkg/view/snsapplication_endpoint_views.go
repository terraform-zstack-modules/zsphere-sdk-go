// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSApplicationEndpointInventoryView SNSApplicationEndpoint
type SNSApplicationEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// UpdateSNSApplicationEndpointEventView UpdateSNSApplicationEndpointEvent
type UpdateSNSApplicationEndpointEventView struct {
	Inventory SNSApplicationEndpointInventoryView `json:"inventory,omitempty"`
}

// DeleteSNSApplicationEndpointEventView DeleteSNSApplicationEndpointEvent
type DeleteSNSApplicationEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeSNSApplicationEndpointStateEventView ChangeSNSApplicationEndpointStateEvent
type ChangeSNSApplicationEndpointStateEventView struct {
	Inventory SNSApplicationEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSApplicationEndpointView QuerySNSApplicationEndpoint
type QuerySNSApplicationEndpointView struct {
	Inventories []SNSApplicationEndpointInventoryView `json:"inventories,omitempty"`
}

