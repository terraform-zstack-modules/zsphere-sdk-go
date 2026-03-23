// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PriceTableInventoryView PriceTable
type PriceTableInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
}

// CreatePriceTableEventView CreatePriceTableEvent
type CreatePriceTableEventView struct {
	Inventory PriceTableInventoryView `json:"inventory,omitempty"`
}

// ChangeAccountPriceTableBindingEventView ChangeAccountPriceTableBindingEvent
type ChangeAccountPriceTableBindingEventView struct {
	Inventory PriceTableInventoryView `json:"inventory,omitempty"`
}

// QueryPriceTableRelyView QueryPriceTableRely
type QueryPriceTableRelyView struct {
	Inventories []PriceTableInventoryView `json:"inventories,omitempty"`
}

// UpdatePriceTableEventView UpdatePriceTableEvent
type UpdatePriceTableEventView struct {
	Inventory PriceTableInventoryView `json:"inventory,omitempty"`
}

// AttachPriceTableToAccountEventView AttachPriceTableToAccountEvent
type AttachPriceTableToAccountEventView struct {
	Inventory PriceTableInventoryView `json:"inventory,omitempty"`
}

// DetachPriceTableFromAccountEventView DetachPriceTableFromAccountEvent
type DetachPriceTableFromAccountEventView struct {
	Inventory PriceTableInventoryView `json:"inventory,omitempty"`
}

// DeletePriceTableEventView DeletePriceTableEvent
type DeletePriceTableEventView struct {
	Success bool `json:"success,omitempty"`
}

