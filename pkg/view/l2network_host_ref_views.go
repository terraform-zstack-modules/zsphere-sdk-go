// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2NetworkHostRefInventoryView L2NetworkHostRef
type L2NetworkHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
	L2ProviderType string `json:"l2ProviderType,omitempty"`
	BridgeName string `json:"bridgeName,omitempty"`
	SkipDeletion bool `json:"skipDeletion,omitempty"`
}

