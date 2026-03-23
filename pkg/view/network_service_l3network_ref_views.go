// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NetworkServiceL3NetworkRefInventoryView NetworkServiceL3NetworkRef
type NetworkServiceL3NetworkRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	NetworkServiceProviderUuid string `json:"networkServiceProviderUuid,omitempty"`
	NetworkServiceType string `json:"networkServiceType,omitempty"`
}

// QueryNetworkServiceL3NetworkRefView QueryNetworkServiceL3NetworkRef
type QueryNetworkServiceL3NetworkRefView struct {
	Inventories []NetworkServiceL3NetworkRefInventoryView `json:"inventories,omitempty"`
}

