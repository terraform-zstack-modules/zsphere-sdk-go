// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DeviceAddressView DeviceAddress
type DeviceAddressView struct {
	Type string `json:"type,omitempty"`
	Bus string `json:"bus,omitempty"`
	Domain string `json:"domain,omitempty"`
	Slot string `json:"slot,omitempty"`
	Function string `json:"function,omitempty"`
	Controller string `json:"controller,omitempty"`
	Target string `json:"target,omitempty"`
	Unit string `json:"unit,omitempty"`
}

