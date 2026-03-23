// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZceXThirdPartyPlatformAlertRefInventoryView ZceXThirdPartyPlatformAlertRef
type ZceXThirdPartyPlatformAlertRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZceXUuid string `json:"zceXUuid,omitempty"`
	ThirdPartyPlatformUuid string `json:"thirdPartyPlatformUuid,omitempty"`
}

// QueryZceXThirdPartyPlatformAlertRefView QueryZceXThirdPartyPlatformAlertRef
type QueryZceXThirdPartyPlatformAlertRefView struct {
	Inventories []ZceXThirdPartyPlatformAlertRefInventoryView `json:"inventories,omitempty"`
}

// CreateZceXAlertPlatformEventView CreateZceXAlertPlatformEvent
type CreateZceXAlertPlatformEventView struct {
	ThirdPartyPlatform ThirdpartyPlatformInventoryView `json:"thirdPartyPlatform,omitempty"`
	Inventory ZceXThirdPartyPlatformAlertRefInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

