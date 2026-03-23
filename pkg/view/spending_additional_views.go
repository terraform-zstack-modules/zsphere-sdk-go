// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SpendingView Spending
type SpendingView struct {
	SpendingType string `json:"spendingType,omitempty"`
	Spending float64 `json:"spending,omitempty"`
	HypervisorTypeSpending map[string]float64 `json:"hypervisorTypeSpending,omitempty"`
	DateStart int64 `json:"dateStart,omitempty"`
	DateEnd int64 `json:"dateEnd,omitempty"`
	Details []SpendingDetailsView `json:"details,omitempty"`
}

