// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SshPrivateKeyPairInventoryView SshPrivateKeyPair
type SshPrivateKeyPairInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
}

// GenerateSshKeyPairView GenerateSshKeyPair
type GenerateSshKeyPairView struct {
	Inventory SshPrivateKeyPairInventoryView `json:"inventory,omitempty"`
}

