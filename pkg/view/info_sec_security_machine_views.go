// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// InfoSecSecurityMachineInventoryView InfoSecSecurityMachine
type InfoSecSecurityMachineInventoryView struct {
	BaseInfoView
	BaseTimeView
	Port int `json:"port,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	SecretResourcePoolUuid string `json:"secretResourcePoolUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Type string `json:"type,omitempty"`
	Model string `json:"model,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
}

// UpdateSecurityMachineEventView UpdateSecurityMachineEvent
type UpdateSecurityMachineEventView struct {
	Inventory SecurityMachineInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddSecurityMachineEventView AddSecurityMachineEvent
type AddSecurityMachineEventView struct {
	Inventory SecurityMachineInventoryView `json:"inventory,omitempty"`
}

