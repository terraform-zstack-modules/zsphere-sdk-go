// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VdiPortInfoView VdiPortInfo
type VdiPortInfoView struct {
	VncPort int `json:"vncPort,omitempty"`
	SpicePort int `json:"spicePort,omitempty"`
	SpiceTlsPort int `json:"spiceTlsPort,omitempty"`
}

