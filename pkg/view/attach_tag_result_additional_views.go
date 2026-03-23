// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AttachTagResultView AttachTagResult
type AttachTagResultView struct {
	Inventory UserTagInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

