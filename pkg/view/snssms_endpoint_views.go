// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSSmsEndpointInventoryView SNSSmsEndpoint
type SNSSmsEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Receivers []SNSSmsReceiverInventoryView `json:"receivers,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// QuerySNSSmsEndpointView QuerySNSSmsEndpoint
type QuerySNSSmsEndpointView struct {
	Inventories []SNSSmsEndpointInventoryView `json:"inventories,omitempty"`
}

