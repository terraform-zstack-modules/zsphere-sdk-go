// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SSORedirectTemplateInventoryView SSORedirectTemplate
type SSORedirectTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ClientUuid string `json:"clientUuid,omitempty"`
	RedirectTemplate string `json:"redirectTemplate,omitempty"`
}

// CreateSSORedirectTemplateEventView CreateSSORedirectTemplateEvent
type CreateSSORedirectTemplateEventView struct {
	Inventory SSORedirectTemplateInventoryView `json:"inventory,omitempty"`
}

// UpdateSSORedirectTemplateEventView UpdateSSORedirectTemplateEvent
type UpdateSSORedirectTemplateEventView struct {
	Inventory SSORedirectTemplateInventoryView `json:"inventory,omitempty"`
}

// DeleteSSORedirectTemplateEventView DeleteSSORedirectTemplateEvent
type DeleteSSORedirectTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

