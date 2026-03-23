// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MetricDataHttpReceiverInventoryView MetricDataHttpReceiver
type MetricDataHttpReceiverInventoryView struct {
	BaseInfoView
	BaseTimeView
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

// DeleteMetricDataHttpReceiverEventView DeleteMetricDataHttpReceiverEvent
type DeleteMetricDataHttpReceiverEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryMetricDataHttpReceiverView QueryMetricDataHttpReceiver
type QueryMetricDataHttpReceiverView struct {
	Inventories []MetricDataHttpReceiverInventoryView `json:"inventories,omitempty"`
}

// CreateMetricDataHttpReceiverEventView CreateMetricDataHttpReceiverEvent
type CreateMetricDataHttpReceiverEventView struct {
	Inventory MetricDataHttpReceiverInventoryView `json:"inventory,omitempty"`
}

