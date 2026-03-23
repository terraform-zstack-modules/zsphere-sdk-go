// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSSnmpPlatformInventoryView SNSSnmpPlatform
type SNSSnmpPlatformInventoryView struct {
	BaseInfoView
	BaseTimeView
	SnmpAddress string `json:"snmpAddress,omitempty"`
	SnmpPort int `json:"snmpPort,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
}

// QuerySNSEmailPlatformView QuerySNSEmailPlatform
type QuerySNSEmailPlatformView struct {
	Inventories []SNSEmailPlatformInventoryView `json:"inventories,omitempty"`
}

