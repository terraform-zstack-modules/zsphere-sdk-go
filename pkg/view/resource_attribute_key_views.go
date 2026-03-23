// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceAttributeKeyInventoryView ResourceAttributeKey
type ResourceAttributeKeyInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	Constraints []ResourceAttributeConstraintInventoryView `json:"constraints,omitempty"`
}

// QueryResourceAttributeKeyView QueryResourceAttributeKey
type QueryResourceAttributeKeyView struct {
	Inventories []ResourceAttributeKeyInventoryView `json:"inventories,omitempty"`
}

// DeleteResourceAttributeKeyEventView DeleteResourceAttributeKeyEvent
type DeleteResourceAttributeKeyEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateResourceAttributeKeyEventView UpdateResourceAttributeKeyEvent
type UpdateResourceAttributeKeyEventView struct {
	Inventory ResourceAttributeKeyInventoryView `json:"inventory,omitempty"`
}

// CreateResourceAttributeKeyEventView CreateResourceAttributeKeyEvent
type CreateResourceAttributeKeyEventView struct {
	Inventory ResourceAttributeKeyInventoryView `json:"inventory,omitempty"`
}

