// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// KvmHypervisorInfoInventoryView KvmHypervisorInfo
type KvmHypervisorInfoInventoryView struct {
	BaseInfoView
	BaseTimeView
	Hypervisor string `json:"hypervisor,omitempty"`
	Version string `json:"version,omitempty"`
	MatchState string `json:"matchState,omitempty"`
}

// QueryKvmHypervisorInfoView QueryKvmHypervisorInfo
type QueryKvmHypervisorInfoView struct {
	Inventories []KvmHypervisorInfoInventoryView `json:"inventories,omitempty"`
}

