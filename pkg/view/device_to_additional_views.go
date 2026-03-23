// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DeviceTOView DeviceTO
type DeviceTOView struct {
	Disk string `json:"disk,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	Target string `json:"target,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

