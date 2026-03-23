// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// QuotaInventoryView Quota
type QuotaInventoryView struct {
	BaseInfoView
	BaseTimeView
	IdentityUuid string `json:"identityUuid,omitempty"`
	IdentityType string `json:"identityType,omitempty"`
	Value int64 `json:"value,omitempty"`
}

// QueryQuotaView QueryQuota
type QueryQuotaView struct {
	Inventories []QuotaInventoryView `json:"inventories,omitempty"`
}

// UpdateQuotaEventView UpdateQuotaEvent
type UpdateQuotaEventView struct {
	Inventory QuotaInventoryView `json:"inventory,omitempty"`
}

