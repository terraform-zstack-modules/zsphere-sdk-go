// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceStackInventoryView ResourceStack
type ResourceStackInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	ParamContent string `json:"paramContent,omitempty"`
	Status string `json:"status,omitempty"`
	Reason string `json:"reason,omitempty"`
	Outputs string `json:"outputs,omitempty"`
	EnableRollback bool `json:"enableRollback,omitempty"`
}

// DeleteResourceStackEventView DeleteResourceStackEvent
type DeleteResourceStackEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateResourceStackEventView CreateResourceStackEvent
type CreateResourceStackEventView struct {
	Inventory ResourceStackInventoryView `json:"inventory,omitempty"`
}

// QueryResourceStackView QueryResourceStack
type QueryResourceStackView struct {
	Inventories []ResourceStackInventoryView `json:"inventories,omitempty"`
}

// UpdateResourceStackEventView UpdateResourceStackEvent
type UpdateResourceStackEventView struct {
	Inventory ResourceStackInventoryView `json:"inventory,omitempty"`
}

