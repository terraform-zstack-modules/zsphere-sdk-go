// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSFeiShuAtPersonInventoryView SNSFeiShuAtPerson
type SNSFeiShuAtPersonInventoryView struct {
	BaseInfoView
	BaseTimeView
	UserId string `json:"userId,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	Remark string `json:"remark,omitempty"`
}

// QuerySNSFeiShuAtPersonView QuerySNSFeiShuAtPerson
type QuerySNSFeiShuAtPersonView struct {
	Inventories []SNSFeiShuAtPersonInventoryView `json:"inventories,omitempty"`
}

// RemoveSNSFeiShuAtPersonEventView RemoveSNSFeiShuAtPersonEvent
type RemoveSNSFeiShuAtPersonEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateAtPersonOfFeiShuEndpointEventView UpdateAtPersonOfFeiShuEndpointEvent
type UpdateAtPersonOfFeiShuEndpointEventView struct {
	Inventory SNSFeiShuAtPersonInventoryView `json:"inventory,omitempty"`
}

// AddSNSFeiShuAtPersonEventView AddSNSFeiShuAtPersonEvent
type AddSNSFeiShuAtPersonEventView struct {
	Inventory SNSFeiShuAtPersonInventoryView `json:"inventory,omitempty"`
}

