// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfCdDriverInfoView OvfCdDriverInfo
type OvfCdDriverInfoView struct {
	AutoAllocation bool `json:"autoAllocation,omitempty"`
	DriverType string `json:"driverType,omitempty"`
	SubType string `json:"subType,omitempty"`
	Name string `json:"name,omitempty"`
}

