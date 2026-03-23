// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SyncTaskResultView SyncTaskResult
type SyncTaskResultView struct {
	SourceUuid string `json:"sourceUuid,omitempty"`
	SourceType string `json:"sourceType,omitempty"`
	ImportStage ImportStageView `json:"importStage,omitempty"`
	CleanStage CleanStageView `json:"cleanStage,omitempty"`
}

