// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CustomPreconfigurationInventoryView CustomPreconfiguration
type CustomPreconfigurationInventoryView struct {
	BaseInfoView
	BaseTimeView
	BaremetalInstanceUuid string `json:"baremetalInstanceUuid,omitempty"`
	Param string `json:"param,omitempty"`
	Value string `json:"value,omitempty"`
}

