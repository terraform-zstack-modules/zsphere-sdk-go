// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EventSubscriptionActionInventoryView EventSubscriptionAction
type EventSubscriptionActionInventoryView struct {
	BaseInfoView
	BaseTimeView
	SubscriptionUuid string `json:"subscriptionUuid,omitempty"`
	ActionType string `json:"actionType,omitempty"`
	ActionUuid string `json:"actionUuid,omitempty"`
}

