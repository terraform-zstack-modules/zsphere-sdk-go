// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSHttpEndpointInventoryView SNSHttpEndpoint
type SNSHttpEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// CreateSNSHttpEndpointEventView CreateSNSHttpEndpointEvent
type CreateSNSHttpEndpointEventView struct {
	Inventory SNSHttpEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSHttpEndpointView QuerySNSHttpEndpoint
type QuerySNSHttpEndpointView struct {
	Inventories []SNSHttpEndpointInventoryView `json:"inventories,omitempty"`
}

