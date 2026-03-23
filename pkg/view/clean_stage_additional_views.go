// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CleanStageView CleanStage
type CleanStageView struct {
	Total int `json:"total,omitempty"`
	Success int `json:"success,omitempty"`
	Skip int `json:"skip,omitempty"`
	Fail int `json:"fail,omitempty"`
}

