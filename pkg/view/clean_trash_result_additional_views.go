// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CleanTrashResultView CleanTrashResult
type CleanTrashResultView struct {
	ResourceUuids []string `json:"resourceUuids,omitempty"`
	Details []string `json:"details,omitempty"`
	Size int64 `json:"size,omitempty"`
}

