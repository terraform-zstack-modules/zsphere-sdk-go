// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSEmailEndpointInventoryView SNSEmailEndpoint
type SNSEmailEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Email string `json:"email,omitempty"`
	EmailAddresses []SNSEmailAddressInventoryView `json:"emailAddresses,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// CreateSNSApplicationEndpointEventView CreateSNSApplicationEndpointEvent
type CreateSNSApplicationEndpointEventView struct {
	Inventory SNSApplicationEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSEmailEndpointView QuerySNSEmailEndpoint
type QuerySNSEmailEndpointView struct {
	Inventories []SNSEmailEndpointInventoryView `json:"inventories,omitempty"`
}

