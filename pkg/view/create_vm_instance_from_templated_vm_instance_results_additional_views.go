// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CreateVmInstanceFromTemplatedVmInstanceResultsView CreateVmInstanceFromTemplatedVmInstanceResults
type CreateVmInstanceFromTemplatedVmInstanceResultsView struct {
	NumberOfClonedVm int `json:"numberOfClonedVm,omitempty"`
	Inventories []CloneVmInstanceInventoryView `json:"inventories,omitempty"`
}

