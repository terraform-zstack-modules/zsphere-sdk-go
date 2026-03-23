// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ElaborationCheckResultView ElaborationCheckResult
type ElaborationCheckResultView struct {
	FileName string `json:"fileName,omitempty"`
	Content string `json:"content,omitempty"`
	Reason string `json:"reason,omitempty"`
}

