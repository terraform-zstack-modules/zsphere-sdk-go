// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FiberChannelStorageInventoryView FiberChannelStorage
type FiberChannelStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Wwnn string `json:"wwnn,omitempty"`
	State string `json:"state,omitempty"`
	FiberChannelLuns []FiberChannelLunInventoryView `json:"fiberChannelLuns,omitempty"`
}

// RefreshFiberChannelStorageEventView RefreshFiberChannelStorageEvent
type RefreshFiberChannelStorageEventView struct {
	Inventories []FiberChannelStorageInventoryView `json:"inventories,omitempty"`
}

// QueryFiberChannelStorageView QueryFiberChannelStorage
type QueryFiberChannelStorageView struct {
	Inventories []FiberChannelStorageInventoryView `json:"inventories,omitempty"`
}

