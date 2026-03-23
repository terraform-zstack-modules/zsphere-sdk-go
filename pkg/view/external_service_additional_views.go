// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ExternalServiceInventoryView ExternalService
type ExternalServiceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Status string `json:"status,omitempty"`
	Capabilities ExternalServiceCapabilitiesView `json:"capabilities,omitempty"`
}

