// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CloneVmInstanceInventoryView CloneVmInstance
type CloneVmInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

