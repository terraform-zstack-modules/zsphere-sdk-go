// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FlowRouterInventoryView FlowRouter
type FlowRouterInventoryView struct {
	BaseInfoView
	BaseTimeView
	SystemID int64 `json:"systemID,omitempty"`
	Type string `json:"type,omitempty"`
}

