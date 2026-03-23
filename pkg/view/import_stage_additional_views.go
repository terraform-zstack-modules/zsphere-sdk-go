// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImportStageView ImportStage
type ImportStageView struct {
	Total int `json:"total,omitempty"`
	Success int `json:"success,omitempty"`
	Fail int `json:"fail,omitempty"`
}

