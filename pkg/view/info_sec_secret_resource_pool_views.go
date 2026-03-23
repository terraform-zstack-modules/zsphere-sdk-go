// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// InfoSecSecretResourcePoolInventoryView InfoSecSecretResourcePool
type InfoSecSecretResourcePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	ConnectionMode int `json:"connectionMode,omitempty"`
	ActivatedToken string `json:"activatedToken,omitempty"`
	ProtectToken string `json:"protectToken,omitempty"`
	HmacToken string `json:"hmacToken,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

