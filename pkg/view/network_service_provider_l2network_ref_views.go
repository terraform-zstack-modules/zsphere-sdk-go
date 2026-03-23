// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NetworkServiceProviderL2NetworkRefInventoryView NetworkServiceProviderL2NetworkRef
type NetworkServiceProviderL2NetworkRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	NetworkServiceProviderUuid string `json:"networkServiceProviderUuid,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
}

