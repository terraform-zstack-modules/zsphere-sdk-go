// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ScsiLunClusterStatusInventoryView ScsiLunClusterStatus
type ScsiLunClusterStatusInventoryView struct {
	BaseInfoView
	BaseTimeView
	AttachedHosts []HostInventoryView `json:"attachedHosts,omitempty"`
	UnattachedHosts []HostInventoryView `json:"unattachedHosts,omitempty"`
	IsAllHostsAttached bool `json:"isAllHostsAttached,omitempty"`
}

