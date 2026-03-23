// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// WebhookInventoryView Webhook
type WebhookInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
	Opaque string `json:"opaque,omitempty"`
}

// CreateWebhookEventView CreateWebhookEvent
type CreateWebhookEventView struct {
	Inventory WebhookInventoryView `json:"inventory,omitempty"`
}

// QueryWebhookView QueryWebhook
type QueryWebhookView struct {
	Inventories []WebhookInventoryView `json:"inventories,omitempty"`
}

// UpdateWebhookEventView UpdateWebhookEvent
type UpdateWebhookEventView struct {
	Inventory WebhookInventoryView `json:"inventory,omitempty"`
}

// DeleteWebhookEventView DeleteWebhookEvent
type DeleteWebhookEventView struct {
	Success bool `json:"success,omitempty"`
}

