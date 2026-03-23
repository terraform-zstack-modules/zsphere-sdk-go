// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CdpPolicyInventoryView CdpPolicy
type CdpPolicyInventoryView struct {
	BaseInfoView
	BaseTimeView
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	RetentionTimePerDay int `json:"retentionTimePerDay,omitempty"`
	HourlyRpSinceDay int `json:"hourlyRpSinceDay,omitempty"`
	DailyRpSinceDay int `json:"dailyRpSinceDay,omitempty"`
	ExpireTimeInDay int `json:"expireTimeInDay,omitempty"`
	FullBackupIntervalInDay int `json:"fullBackupIntervalInDay,omitempty"`
	RecoveryPointPerSecond int `json:"recoveryPointPerSecond,omitempty"`
}

// DeleteCdpPolicyEventView DeleteCdpPolicyEvent
type DeleteCdpPolicyEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateCdpPolicyEventView CreateCdpPolicyEvent
type CreateCdpPolicyEventView struct {
	Inventory CdpPolicyInventoryView `json:"inventory,omitempty"`
}

// QueryCdpPolicyView QueryCdpPolicy
type QueryCdpPolicyView struct {
	Inventories []CdpPolicyInventoryView `json:"inventories,omitempty"`
}

// UpdateCdpPolicyEventView UpdateCdpPolicyEvent
type UpdateCdpPolicyEventView struct {
	Inventory CdpPolicyInventoryView `json:"inventory,omitempty"`
}

