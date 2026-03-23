// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// StackTemplateInventoryView StackTemplate
type StackTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
	State bool `json:"state,omitempty"`
	Content string `json:"content,omitempty"`
	Md5sum string `json:"md5sum,omitempty"`
}

// DeleteStackTemplateEventView DeleteStackTemplateEvent
type DeleteStackTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryStackTemplateView QueryStackTemplate
type QueryStackTemplateView struct {
	Inventories []StackTemplateInventoryView `json:"inventories,omitempty"`
}

// UpdateStackTemplateEventView UpdateStackTemplateEvent
type UpdateStackTemplateEventView struct {
	Inventory StackTemplateInventoryView `json:"inventory,omitempty"`
}

// AddStackTemplateEventView AddStackTemplateEvent
type AddStackTemplateEventView struct {
	Inventory StackTemplateInventoryView `json:"inventory,omitempty"`
}

