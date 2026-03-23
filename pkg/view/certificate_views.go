// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CertificateInventoryView Certificate
type CertificateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Certificate string `json:"certificate,omitempty"`
	Description string `json:"description,omitempty"`
	Listeners []LoadBalancerListenerCertificateRefInventoryView `json:"listeners,omitempty"`
}

// DeleteCertificateEventView DeleteCertificateEvent
type DeleteCertificateEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateCertificateEventView UpdateCertificateEvent
type UpdateCertificateEventView struct {
	Inventory CertificateInventoryView `json:"inventory,omitempty"`
}

// CreateCertificateEventView CreateCertificateEvent
type CreateCertificateEventView struct {
	Inventory CertificateInventoryView `json:"inventory,omitempty"`
}

// QueryCertificateView QueryCertificate
type QueryCertificateView struct {
	Inventories []CertificateInventoryView `json:"inventories,omitempty"`
}

