// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VipNetworkServicesRefInventoryView VipNetworkServicesRef
type VipNetworkServicesRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ServiceType string `json:"serviceType,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
}

