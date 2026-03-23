// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NetworkServiceProviderInventoryView NetworkServiceProvider
type NetworkServiceProviderInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	NetworkServiceTypes []string `json:"networkServiceTypes,omitempty"`
	AttachedL2NetworkUuids []string `json:"attachedL2NetworkUuids,omitempty"`
}

// QueryNetworkServiceProviderView QueryNetworkServiceProvider
type QueryNetworkServiceProviderView struct {
	Inventories []NetworkServiceProviderInventoryView `json:"inventories,omitempty"`
}

