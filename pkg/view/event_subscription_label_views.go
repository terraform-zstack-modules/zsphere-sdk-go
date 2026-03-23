// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EventSubscriptionLabelInventoryView EventSubscriptionLabel
type EventSubscriptionLabelInventoryView struct {
	BaseInfoView
	BaseTimeView
	Key string `json:"key,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value string `json:"value,omitempty"`
}

// UpdateEventSubscriptionLabelEventView UpdateEventSubscriptionLabelEvent
type UpdateEventSubscriptionLabelEventView struct {
	Inventory EventSubscriptionLabelInventoryView `json:"inventory,omitempty"`
}

// AddLabelToEventSubscriptionEventView AddLabelToEventSubscriptionEvent
type AddLabelToEventSubscriptionEventView struct {
	Inventory EventSubscriptionLabelInventoryView `json:"inventory,omitempty"`
}

