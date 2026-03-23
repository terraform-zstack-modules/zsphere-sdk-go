// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfSystemInfoView OvfSystemInfo
type OvfSystemInfoView struct {
	VirtualSystemType string `json:"virtualSystemType,omitempty"`
	FirmwareType string `json:"firmwareType,omitempty"`
}

