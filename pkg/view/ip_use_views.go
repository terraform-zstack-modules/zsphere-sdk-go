// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IpUseInventoryView IpUse
type IpUseInventoryView struct {
	BaseInfoView
	BaseTimeView
	UsedIpUuid string `json:"usedIpUuid,omitempty"`
	ServiceId string `json:"serviceId,omitempty"`
	Use string `json:"use,omitempty"`
	Details string `json:"details,omitempty"`
}

