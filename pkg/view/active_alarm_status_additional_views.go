// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ActiveAlarmStatusView ActiveAlarmStatus
type ActiveAlarmStatusView struct {
	Namespace string `json:"namespace,omitempty"`
	Status string `json:"status,omitempty"`
}

