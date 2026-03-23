// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CreateResourceAttributeResultView CreateResourceAttributeResult
type CreateResourceAttributeResultView struct {
	Inventory ResourceAttributeValueInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

