// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfEthernetAdapterInfoView OvfEthernetAdapterInfo
type OvfEthernetAdapterInfoView struct {
	NetworkName string `json:"networkName,omitempty"`
	NicModel string `json:"nicModel,omitempty"`
	NicName string `json:"nicName,omitempty"`
	AutoAllocation bool `json:"autoAllocation,omitempty"`
}

