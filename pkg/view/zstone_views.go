// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZStoneInventoryView ZStone
type ZStoneInventoryView struct {
	BaseInfoView
	BaseTimeView
	Username string `json:"username,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	AuthorizationServer string `json:"authorizationServer,omitempty"`
	LogInPort int `json:"logInPort,omitempty"`
	ApiPort int `json:"apiPort,omitempty"`
	LogInUrl string `json:"logInUrl,omitempty"`
}

// AddZStoneEventView AddZStoneEvent
type AddZStoneEventView struct {
	Inventory ZStoneInventoryView `json:"inventory,omitempty"`
}

// QueryZStoneView QueryZStone
type QueryZStoneView struct {
	Inventories []ZStoneInventoryView `json:"inventories,omitempty"`
}

// RemoveZStoneEventView RemoveZStoneEvent
type RemoveZStoneEventView struct {
	Success bool `json:"success,omitempty"`
}

