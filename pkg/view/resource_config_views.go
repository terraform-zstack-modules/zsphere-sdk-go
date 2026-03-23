// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceConfigInventoryView ResourceConfig
type ResourceConfigInventoryView struct {
	BaseInfoView
	BaseTimeView
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Description string `json:"description,omitempty"`
	Category string `json:"category,omitempty"`
	Value string `json:"value,omitempty"`
}

// QueryResourceConfigView QueryResourceConfig
type QueryResourceConfigView struct {
	Inventories []ResourceConfigInventoryView `json:"inventories,omitempty"`
}

// GetResourceConfigView GetResourceConfig
type GetResourceConfigView struct {
	Value string `json:"value,omitempty"`
	EffectiveConfigs []ResourceConfigInventoryView `json:"effectiveConfigs,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UpdateResourceConfigEventView UpdateResourceConfigEvent
type UpdateResourceConfigEventView struct {
	Inventory ResourceConfigInventoryView `json:"inventory,omitempty"`
}

// DeleteResourceConfigEventView DeleteResourceConfigEvent
type DeleteResourceConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

