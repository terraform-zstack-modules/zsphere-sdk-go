// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SshKeyPairInventoryView SshKeyPair
type SshKeyPairInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

// CreateSshKeyPairEventView CreateSshKeyPairEvent
type CreateSshKeyPairEventView struct {
	Inventory SshKeyPairInventoryView `json:"inventory,omitempty"`
}

// UpdateSshKeyPairEventView UpdateSshKeyPairEvent
type UpdateSshKeyPairEventView struct {
	Inventory SshKeyPairInventoryView `json:"inventory,omitempty"`
}

// QuerySshKeyPairView QuerySshKeyPair
type QuerySshKeyPairView struct {
	Inventories []SshKeyPairInventoryView `json:"inventories,omitempty"`
}

// AttachSshKeyPairToVmInstanceEventView AttachSshKeyPairToVmInstanceEvent
type AttachSshKeyPairToVmInstanceEventView struct {
	Inventory SshKeyPairInventoryView `json:"inventory,omitempty"`
}

// DeleteSshKeyPairEventView DeleteSshKeyPairEvent
type DeleteSshKeyPairEventView struct {
	Success bool `json:"success,omitempty"`
}

