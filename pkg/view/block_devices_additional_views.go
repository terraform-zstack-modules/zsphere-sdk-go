// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BlockDevicesView BlockDevices
type BlockDevicesView struct {
	UnusedBlockDevices []BlockDeviceView `json:"unusedBlockDevices,omitempty"`
	UsedBlockDevices []BlockDeviceView `json:"usedBlockDevices,omitempty"`
}

