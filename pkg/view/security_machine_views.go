// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SecurityMachineInventoryView SecurityMachine
type SecurityMachineInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZoneUuid string `json:"zoneUuid,omitempty"`
	SecretResourcePoolUuid string `json:"secretResourcePoolUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Type string `json:"type,omitempty"`
	Model string `json:"model,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
}

// SetSecurityMachineKeyEventView SetSecurityMachineKeyEvent
type SetSecurityMachineKeyEventView struct {
	Inventories []SecurityMachineInventoryView `json:"inventories,omitempty"`
}

// ChangeSecurityMachineStateEventView ChangeSecurityMachineStateEvent
type ChangeSecurityMachineStateEventView struct {
	Inventory SecurityMachineInventoryView `json:"inventory,omitempty"`
}

// DeleteSecurityMachineEventView DeleteSecurityMachineEvent
type DeleteSecurityMachineEventView struct {
	Success bool `json:"success,omitempty"`
}

// QuerySecurityMachineView QuerySecurityMachine
type QuerySecurityMachineView struct {
	Inventories []SecurityMachineInventoryView `json:"inventories,omitempty"`
}

