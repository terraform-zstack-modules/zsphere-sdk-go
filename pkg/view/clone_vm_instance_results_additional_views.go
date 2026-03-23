// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CloneVmInstanceResultsView CloneVmInstanceResults
type CloneVmInstanceResultsView struct {
	NumberOfClonedVm int `json:"numberOfClonedVm,omitempty"`
	Inventories []CloneVmInstanceInventoryView `json:"inventories,omitempty"`
}

