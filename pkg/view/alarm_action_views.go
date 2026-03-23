// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlarmActionInventoryView AlarmAction
type AlarmActionInventoryView struct {
	BaseInfoView
	BaseTimeView
	AlarmUuid string `json:"alarmUuid,omitempty"`
	ActionType string `json:"actionType,omitempty"`
	ActionUuid string `json:"actionUuid,omitempty"`
}

