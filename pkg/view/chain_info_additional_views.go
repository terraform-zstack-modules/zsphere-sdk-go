// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ChainInfoView ChainInfo
type ChainInfoView struct {
	RunningTask []RunningTaskInfoView `json:"runningTask,omitempty"`
	PendingTask []PendingTaskInfoView `json:"pendingTask,omitempty"`
}

