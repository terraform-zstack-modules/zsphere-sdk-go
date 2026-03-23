// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSDingTalkEndpointInventoryView SNSDingTalkEndpoint
type SNSDingTalkEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	Secret string `json:"secret,omitempty"`
	AtPersonPhoneNumbers []string `json:"atPersonPhoneNumbers,omitempty"`
	AtPersonList []SNSDingTalkAtPersonInventoryView `json:"atPersonList,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// CreateSNSDingTalkEndpointEventView CreateSNSDingTalkEndpointEvent
type CreateSNSDingTalkEndpointEventView struct {
	Inventory SNSDingTalkEndpointInventoryView `json:"inventory,omitempty"`
}

// QuerySNSDingTalkEndpointView QuerySNSDingTalkEndpoint
type QuerySNSDingTalkEndpointView struct {
	Inventories []SNSDingTalkEndpointInventoryView `json:"inventories,omitempty"`
}

