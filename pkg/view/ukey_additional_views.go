// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// UKeyInventoryView UKey
type UKeyInventoryView struct {
	BaseInfoView
	BaseTimeView
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	Status string `json:"status,omitempty"`
	KeyId string `json:"keyId,omitempty"`
}

