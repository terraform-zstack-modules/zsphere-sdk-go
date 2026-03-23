// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmCapabilitiesView VmCapabilities
type VmCapabilitiesView struct {
	SupportLiveMigration bool `json:"supportLiveMigration,omitempty"`
	SupportVolumeMigration bool `json:"supportVolumeMigration,omitempty"`
	SupportReimage bool `json:"supportReimage,omitempty"`
	SupportMemorySnapshot bool `json:"supportMemorySnapshot,omitempty"`
}

