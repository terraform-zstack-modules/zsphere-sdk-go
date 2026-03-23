// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FlowMeterInventoryView FlowMeter
type FlowMeterInventoryView struct {
	BaseInfoView
	BaseTimeView
	Collectors []FlowCollectorInventoryView `json:"collectors,omitempty"`
	NetworkRefs []NetworkRouterFlowMeterRefInventoryView `json:"networkRefs,omitempty"`
	Description string `json:"description,omitempty"`
	Sample int64 `json:"sample,omitempty"`
	ExpireInterval int64 `json:"expireInterval,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryFlowMeterView QueryFlowMeter
type QueryFlowMeterView struct {
	Inventories []FlowMeterInventoryView `json:"inventories,omitempty"`
}

// GetVpcAttachedNetflowView GetVpcAttachedNetflow
type GetVpcAttachedNetflowView struct {
	Inventories []FlowMeterInventoryView `json:"inventories,omitempty"`
}

// CreateFlowMeterEventView CreateFlowMeterEvent
type CreateFlowMeterEventView struct {
	Inventory FlowMeterInventoryView `json:"inventory,omitempty"`
}

// DeleteFlowMeterEventView DeleteFlowMeterEvent
type DeleteFlowMeterEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateFlowMeterEventView UpdateFlowMeterEvent
type UpdateFlowMeterEventView struct {
	Inventory FlowMeterInventoryView `json:"inventory,omitempty"`
}

