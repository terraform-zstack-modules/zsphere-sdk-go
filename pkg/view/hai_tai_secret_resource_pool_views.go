// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HaiTaiSecretResourcePoolInventoryView HaiTaiSecretResourcePool
type HaiTaiSecretResourcePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	ManagementIp string `json:"managementIp,omitempty"`
	Port int `json:"port,omitempty"`
	Realm string `json:"realm,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

