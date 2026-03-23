// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSMicrosoftTeamsEndpointInventoryView SNSMicrosoftTeamsEndpoint
type SNSMicrosoftTeamsEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// CreateSNSMicrosoftTeamsEndpointEventView CreateSNSMicrosoftTeamsEndpointEvent
type CreateSNSMicrosoftTeamsEndpointEventView struct {
	Inventory SNSMicrosoftTeamsEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSMicrosoftTeamsEndpointView QuerySNSMicrosoftTeamsEndpoint
type QuerySNSMicrosoftTeamsEndpointView struct {
	Inventories []SNSMicrosoftTeamsEndpointInventoryView `json:"inventories,omitempty"`
}

