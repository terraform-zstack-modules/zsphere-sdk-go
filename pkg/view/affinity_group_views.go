// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AffinityGroupInventoryView AffinityGroup
type AffinityGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Policy string `json:"policy,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	Appliance string `json:"appliance,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	State string `json:"state,omitempty"`
	Usages []AffinityGroupUsageInventoryView `json:"usages,omitempty"`
}

// UpdateAffinityGroupEventView UpdateAffinityGroupEvent
type UpdateAffinityGroupEventView struct {
	Inventory AffinityGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteAffinityGroupEventView DeleteAffinityGroupEvent
type DeleteAffinityGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveVmFromAffinityGroupEventView RemoveVmFromAffinityGroupEvent
type RemoveVmFromAffinityGroupEventView struct {
	Inventory AffinityGroupInventoryView `json:"inventory,omitempty"`
}

// GetCandidateAffinityGroupForCreatingVmView GetCandidateAffinityGroupForCreatingVm
type GetCandidateAffinityGroupForCreatingVmView struct {
	Inventories []AffinityGroupInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ChangeAffinityGroupStateEventView ChangeAffinityGroupStateEvent
type ChangeAffinityGroupStateEventView struct {
	Inventory AffinityGroupInventoryView `json:"inventory,omitempty"`
}

// GetCandidateAffinityGroupForAttachingVmView GetCandidateAffinityGroupForAttachingVm
type GetCandidateAffinityGroupForAttachingVmView struct {
	Inventories []AffinityGroupInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddVmToAffinityGroupEventView AddVmToAffinityGroupEvent
type AddVmToAffinityGroupEventView struct {
	Inventory AffinityGroupInventoryView `json:"inventory,omitempty"`
}

// CreateAffinityGroupEventView CreateAffinityGroupEvent
type CreateAffinityGroupEventView struct {
	Inventory AffinityGroupInventoryView `json:"inventory,omitempty"`
}

// QueryAffinityGroupView QueryAffinityGroup
type QueryAffinityGroupView struct {
	Inventories []AffinityGroupInventoryView `json:"inventories,omitempty"`
}

