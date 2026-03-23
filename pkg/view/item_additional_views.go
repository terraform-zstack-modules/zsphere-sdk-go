// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ItemInventoryView Item
type ItemInventoryView struct {
	BaseInfoView
	BaseTimeView
	ReadableName string `json:"readableName,omitempty"`
}

