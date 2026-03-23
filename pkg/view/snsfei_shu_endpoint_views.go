// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSFeiShuEndpointInventoryView SNSFeiShuEndpoint
type SNSFeiShuEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	AtPersonList []SNSFeiShuAtPersonInventoryView `json:"atPersonList,omitempty"`
	Secret string `json:"secret,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// QuerySNSFeiShuEndpointView QuerySNSFeiShuEndpoint
type QuerySNSFeiShuEndpointView struct {
	Inventories []SNSFeiShuEndpointInventoryView `json:"inventories,omitempty"`
}

// CreateSNSFeiShuEndpointEventView CreateSNSFeiShuEndpointEvent
type CreateSNSFeiShuEndpointEventView struct {
	Inventory SNSFeiShuEndpointInventoryView `json:"inventory,omitempty"`
}

