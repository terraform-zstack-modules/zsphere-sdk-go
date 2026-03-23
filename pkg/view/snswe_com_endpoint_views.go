// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSWeComEndpointInventoryView SNSWeComEndpoint
type SNSWeComEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	AtPersonList []SNSWeComAtPersonInventoryView `json:"atPersonList,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// QuerySNSWeComEndpointView QuerySNSWeComEndpoint
type QuerySNSWeComEndpointView struct {
	Inventories []SNSWeComEndpointInventoryView `json:"inventories,omitempty"`
}

// CreateSNSWeComEndpointEventView CreateSNSWeComEndpointEvent
type CreateSNSWeComEndpointEventView struct {
	Inventory SNSWeComEndpointInventoryView `json:"inventory,omitempty"`
}

