// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostHaStateInventoryView HostHaState
type HostHaStateInventoryView struct {
	BaseInfoView
	BaseTimeView
	State string `json:"state,omitempty"`
}

