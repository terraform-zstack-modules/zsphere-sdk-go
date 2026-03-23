// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZSha2StatusViewView ZSha2StatusView
type ZSha2StatusViewView struct {
	Vip string `json:"vip,omitempty"`
	UiHttpPath string `json:"uiHttpPath,omitempty"`
	Nodes []ManagementNodeStatusViewView `json:"nodes,omitempty"`
}

