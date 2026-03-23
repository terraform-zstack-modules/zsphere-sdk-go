// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ShrinkResultView ShrinkResult
type ShrinkResultView struct {
	OldSize int64 `json:"oldSize,omitempty"`
	Size int64 `json:"size,omitempty"`
	DeltaSize int64 `json:"deltaSize,omitempty"`
}

