// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SkippedResourcesView SkippedResources
type SkippedResourcesView struct {
	VmInstances []VmInstanceInventoryView `json:"vmInstances,omitempty"`
}

