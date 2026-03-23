// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfVolumeInfoView OvfVolumeInfo
type OvfVolumeInfoView struct {
	Name string `json:"name,omitempty"`
	DiskId string `json:"diskId,omitempty"`
	DriverType string `json:"driverType,omitempty"`
}

