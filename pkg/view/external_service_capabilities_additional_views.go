// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ExternalServiceCapabilitiesView ExternalServiceCapabilities
type ExternalServiceCapabilitiesView struct {
	ReloadConfig bool `json:"reloadConfig,omitempty"`
}

