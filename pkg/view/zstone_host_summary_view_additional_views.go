// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZStoneHostSummaryViewView ZStoneHostSummaryView
type ZStoneHostSummaryViewView struct {
	Count int `json:"count,omitempty"`
	AdminIps []string `json:"adminIps,omitempty"`
}

