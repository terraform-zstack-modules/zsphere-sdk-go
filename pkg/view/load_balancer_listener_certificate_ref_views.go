// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoadBalancerListenerCertificateRefInventoryView LoadBalancerListenerCertificateRef
type LoadBalancerListenerCertificateRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	ListenerUuid string `json:"listenerUuid,omitempty"`
	CertificateUuid string `json:"certificateUuid,omitempty"`
}

