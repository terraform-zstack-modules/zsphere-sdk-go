// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NetworkServiceTypeInventoryView NetworkServiceType
type NetworkServiceTypeInventoryView struct {
	BaseInfoView
	BaseTimeView
	NetworkServiceProviderUuid string `json:"networkServiceProviderUuid,omitempty"`
	Type string `json:"type,omitempty"`
}

