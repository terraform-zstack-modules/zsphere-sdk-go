// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// GlobalConfigInventoryView GlobalConfig
type GlobalConfigInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Value string `json:"value,omitempty"`
}

// UpdateGlobalConfigEventView UpdateGlobalConfigEvent
type UpdateGlobalConfigEventView struct {
	Inventory GlobalConfigInventoryView `json:"inventory,omitempty"`
}

// ResetGlobalConfigEventView ResetGlobalConfigEvent
type ResetGlobalConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryGlobalConfigView QueryGlobalConfig
type QueryGlobalConfigView struct {
	Inventories []GlobalConfigInventoryView `json:"inventories,omitempty"`
}

