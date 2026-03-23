// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CephPrimaryStorageMonInventoryView CephPrimaryStorageMon
type CephPrimaryStorageMonInventoryView struct {
	BaseInfoView
	BaseTimeView
	Hostname string `json:"hostname,omitempty"`
	MonPort int `json:"monPort,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	MonAddr string `json:"monAddr,omitempty"`
	SshUsername string `json:"sshUsername,omitempty"`
	SshPassword string `json:"sshPassword,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Status string `json:"status,omitempty"`
	MonUuid string `json:"monUuid,omitempty"`
}

// UpdateCephPrimaryStorageMonEventView UpdateCephPrimaryStorageMonEvent
type UpdateCephPrimaryStorageMonEventView struct {
	Inventory CephPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

