// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSDingTalkAtPersonInventoryView SNSDingTalkAtPerson
type SNSDingTalkAtPersonInventoryView struct {
	BaseInfoView
	BaseTimeView
	PhoneNumber string `json:"phoneNumber,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	Remark string `json:"remark,omitempty"`
}

// RemoveSNSDingTalkAtPersonEventView RemoveSNSDingTalkAtPersonEvent
type RemoveSNSDingTalkAtPersonEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateAtPersonOfDingTalkEndpointEventView UpdateAtPersonOfDingTalkEndpointEvent
type UpdateAtPersonOfDingTalkEndpointEventView struct {
	Inventory SNSDingTalkAtPersonInventoryView `json:"inventory,omitempty"`
}

// AddSNSDingTalkAtPersonEventView AddSNSDingTalkAtPersonEvent
type AddSNSDingTalkAtPersonEventView struct {
	Inventory SNSDingTalkAtPersonInventoryView `json:"inventory,omitempty"`
}

// QuerySNSDingTalkAtPersonView QuerySNSDingTalkAtPerson
type QuerySNSDingTalkAtPersonView struct {
	Inventories []SNSDingTalkAtPersonInventoryView `json:"inventories,omitempty"`
}

