// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HbaDeviceInventoryView HbaDevice
type HbaDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	HbaType string `json:"hbaType,omitempty"`
}

