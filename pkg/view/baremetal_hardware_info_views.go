// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BaremetalHardwareInfoInventoryView BaremetalHardwareInfo
type BaremetalHardwareInfoInventoryView struct {
	BaseInfoView
	BaseTimeView
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

