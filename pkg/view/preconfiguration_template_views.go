// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PreconfigurationTemplateInventoryView PreconfigurationTemplate
type PreconfigurationTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Type string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	Md5sum string `json:"md5sum,omitempty"`
	IsPredefined bool `json:"isPredefined,omitempty"`
	State string `json:"state,omitempty"`
	CustomParams []string `json:"customParams,omitempty"`
}

// ChangePreconfigurationTemplateStateEventView ChangePreconfigurationTemplateStateEvent
type ChangePreconfigurationTemplateStateEventView struct {
	Inventory PreconfigurationTemplateInventoryView `json:"inventory,omitempty"`
}

// DeletePreconfigurationTemplateEventView DeletePreconfigurationTemplateEvent
type DeletePreconfigurationTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdatePreconfigurationTemplateEventView UpdatePreconfigurationTemplateEvent
type UpdatePreconfigurationTemplateEventView struct {
	Inventory PreconfigurationTemplateInventoryView `json:"inventory,omitempty"`
}

// QueryPreconfigurationTemplatesView QueryPreconfigurationTemplates
type QueryPreconfigurationTemplatesView struct {
	Inventories []PreconfigurationTemplateInventoryView `json:"inventories,omitempty"`
}

// AddPreconfigurationTemplateEventView AddPreconfigurationTemplateEvent
type AddPreconfigurationTemplateEventView struct {
	Inventory PreconfigurationTemplateInventoryView `json:"inventory,omitempty"`
}

