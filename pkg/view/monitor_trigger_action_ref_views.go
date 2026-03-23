// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorTriggerActionRefInventoryView MonitorTriggerActionRef
type MonitorTriggerActionRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	TriggerUuid string `json:"triggerUuid,omitempty"`
	ActionUuid string `json:"actionUuid,omitempty"`
}

