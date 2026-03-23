// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSSmsReceiverInventoryView SNSSmsReceiver
type SNSSmsReceiverInventoryView struct {
	BaseInfoView
	BaseTimeView
	PhoneNumber string `json:"phoneNumber,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// AddSNSSmsReceiverEventView AddSNSSmsReceiverEvent
type AddSNSSmsReceiverEventView struct {
	Inventory SNSSmsReceiverInventoryView `json:"inventory,omitempty"`
}

// RemoveSNSSmsReceiverEventView RemoveSNSSmsReceiverEvent
type RemoveSNSSmsReceiverEventView struct {
	Success bool `json:"success,omitempty"`
}

