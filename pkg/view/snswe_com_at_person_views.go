// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSWeComAtPersonInventoryView SNSWeComAtPerson
type SNSWeComAtPersonInventoryView struct {
	BaseInfoView
	BaseTimeView
	UserId string `json:"userId,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	Remark string `json:"remark,omitempty"`
}

// RemoveSNSWeComAtPersonEventView RemoveSNSWeComAtPersonEvent
type RemoveSNSWeComAtPersonEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddSNSWeComAtPersonEventView AddSNSWeComAtPersonEvent
type AddSNSWeComAtPersonEventView struct {
	Inventory SNSWeComAtPersonInventoryView `json:"inventory,omitempty"`
}

// QuerySNSWeComAtPersonView QuerySNSWeComAtPerson
type QuerySNSWeComAtPersonView struct {
	Inventories []SNSWeComAtPersonInventoryView `json:"inventories,omitempty"`
}

// UpdateAtPersonOfWeComEndpointEventView UpdateAtPersonOfWeComEndpointEvent
type UpdateAtPersonOfWeComEndpointEventView struct {
	Inventory SNSWeComAtPersonInventoryView `json:"inventory,omitempty"`
}

