// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TrashCleanupResultView TrashCleanupResult
type TrashCleanupResultView struct {
	ResourceUuid string `json:"resourceUuid,omitempty"`
	Success bool `json:"success,omitempty"`
	TrashId int64 `json:"trashId,omitempty"`
	Size int64 `json:"size,omitempty"`
}

