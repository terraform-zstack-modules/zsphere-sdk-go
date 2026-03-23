// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SpendingDetailsView SpendingDetails
type SpendingDetailsView struct {
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	Spending float64 `json:"spending,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	Type string `json:"type,omitempty"`
}

