// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CCSCertificateAccountRefInventoryView CCSCertificateAccountRef
type CCSCertificateAccountRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountUuid string `json:"accountUuid,omitempty"`
	CertificateUuid string `json:"certificateUuid,omitempty"`
	State string `json:"state,omitempty"`
}

