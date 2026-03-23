// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostKernelInterfaceResultView HostKernelInterfaceResult
type HostKernelInterfaceResultView struct {
	Inventory HostKernelInterfaceInventoryView `json:"inventory,omitempty"`
}

