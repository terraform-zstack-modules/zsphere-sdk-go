// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FlowCollectorInventoryView FlowCollector
type FlowCollectorInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	FlowMeterUuid string `json:"flowMeterUuid,omitempty"`
	Server string `json:"server,omitempty"`
	Port int64 `json:"port,omitempty"`
}

// DeleteFlowCollectorEventView DeleteFlowCollectorEvent
type DeleteFlowCollectorEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateFlowCollectorEventView CreateFlowCollectorEvent
type CreateFlowCollectorEventView struct {
	Inventory FlowCollectorInventoryView `json:"inventory,omitempty"`
}

// QueryFlowCollectorView QueryFlowCollector
type QueryFlowCollectorView struct {
	Inventories []FlowCollectorInventoryView `json:"inventories,omitempty"`
}

