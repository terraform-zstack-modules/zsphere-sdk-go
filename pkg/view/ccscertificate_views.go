// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CCSCertificateInventoryView CCSCertificate
type CCSCertificateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Algorithm string `json:"algorithm,omitempty"`
	Format string `json:"format,omitempty"`
	IssuerDN string `json:"issuerDN,omitempty"`
	SubjectDN string `json:"subjectDN,omitempty"`
	SerNumber string `json:"serNumber,omitempty"`
	EffectiveTime time.Time `json:"effectiveTime,omitempty"`
	ExpirationTime time.Time `json:"expirationTime,omitempty"`
	AccountCertificateRefs []CCSCertificateAccountRefInventoryView `json:"accountCertificateRefs,omitempty"`
}

// AddCCSCertificateEventView AddCCSCertificateEvent
type AddCCSCertificateEventView struct {
	Inventory CCSCertificateInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UpdateCCSCertificateAccountStateEventView UpdateCCSCertificateAccountStateEvent
type UpdateCCSCertificateAccountStateEventView struct {
	Inventory CCSCertificateInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AttachCCSCertificateToAccountEventView AttachCCSCertificateToAccountEvent
type AttachCCSCertificateToAccountEventView struct {
	Inventory CCSCertificateInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteCCSCertificateEventView DeleteCCSCertificateEvent
type DeleteCCSCertificateEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryCCSCertificateView QueryCCSCertificate
type QueryCCSCertificateView struct {
	Inventories []CCSCertificateInventoryView `json:"inventories,omitempty"`
}

