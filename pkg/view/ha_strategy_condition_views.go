// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HaStrategyConditionInventoryView HaStrategyCondition
type HaStrategyConditionInventoryView struct {
	BaseInfoView
	BaseTimeView
	FencerName string `json:"fencerName,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateHaStrategyConditionEventView UpdateHaStrategyConditionEvent
type UpdateHaStrategyConditionEventView struct {
	Inventory HaStrategyConditionInventoryView `json:"inventory,omitempty"`
}

