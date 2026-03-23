// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MttyDeviceInventoryView MttyDevice
type MttyDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	VirtStatus string `json:"virtStatus,omitempty"`
}

// QueryMttyDeviceView QueryMttyDevice
type QueryMttyDeviceView struct {
	Inventories []MttyDeviceInventoryView `json:"inventories,omitempty"`
}

