// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlertInventoryView Alert
type AlertInventoryView struct {
	BaseInfoView
	BaseTimeView
	TriggerUuid string `json:"triggerUuid,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	Content string `json:"content,omitempty"`
}

// DeleteAlertEventView DeleteAlertEvent
type DeleteAlertEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryAlertView QueryAlert
type QueryAlertView struct {
	Inventories []AlertInventoryView `json:"inventories,omitempty"`
}

