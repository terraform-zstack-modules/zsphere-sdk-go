// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImageStoreGcResultView ImageStoreGcResult
type ImageStoreGcResultView struct {
	FreedSpaceInBytes int64 `json:"freedSpaceInBytes,omitempty"`
}

