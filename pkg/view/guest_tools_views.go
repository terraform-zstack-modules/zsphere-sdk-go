// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// GuestToolsInventoryView GuestTools
type GuestToolsInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	Version string `json:"version,omitempty"`
	AgentType string `json:"agentType,omitempty"`
}

// GetLatestGuestToolsForVmView GetLatestGuestToolsForVm
type GetLatestGuestToolsForVmView struct {
	Inventory GuestToolsInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

