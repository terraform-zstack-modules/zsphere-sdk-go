// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EventSubscriptionInventoryView EventSubscription
type EventSubscriptionInventoryView struct {
	BaseInfoView
	BaseTimeView
	Namespace string `json:"namespace,omitempty"`
	EventName string `json:"eventName,omitempty"`
	State string `json:"state,omitempty"`
	Actions []EventSubscriptionActionInventoryView `json:"actions,omitempty"`
	Labels []EventSubscriptionLabelInventoryView `json:"labels,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
}

// ChangeEventSubscriptionStateEventView ChangeEventSubscriptionStateEvent
type ChangeEventSubscriptionStateEventView struct {
	Inventory EventSubscriptionInventoryView `json:"inventory,omitempty"`
}

// AddActionToEventSubscriptionEventView AddActionToEventSubscriptionEvent
type AddActionToEventSubscriptionEventView struct {
	Inventory EventSubscriptionInventoryView `json:"inventory,omitempty"`
}

// QueryEventSubscriptionView QueryEventSubscription
type QueryEventSubscriptionView struct {
	Inventories []EventSubscriptionInventoryView `json:"inventories,omitempty"`
}

// UpdateSubscribeEventEventView UpdateSubscribeEventEvent
type UpdateSubscribeEventEventView struct {
	Inventory EventSubscriptionInventoryView `json:"inventory,omitempty"`
}

