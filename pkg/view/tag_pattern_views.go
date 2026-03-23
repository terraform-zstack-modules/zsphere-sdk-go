// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TagPatternInventoryView TagPattern
type TagPatternInventoryView struct {
	BaseInfoView
	BaseTimeView
	Value string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Color string `json:"color,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryTagView QueryTag
type QueryTagView struct {
	Inventories []TagPatternInventoryView `json:"inventories,omitempty"`
}

// CreateTagEventView CreateTagEvent
type CreateTagEventView struct {
	Inventory TagPatternInventoryView `json:"inventory,omitempty"`
}

// UpdateTagEventView UpdateTagEvent
type UpdateTagEventView struct {
	Inventory TagPatternInventoryView `json:"inventory,omitempty"`
}

