// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSTextTemplateInventoryView SNSTextTemplate
type SNSTextTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ApplicationPlatformType string `json:"applicationPlatformType,omitempty"`
	Subject string `json:"subject,omitempty"`
	RecoverySubject string `json:"recoverySubject,omitempty"`
	Template string `json:"template,omitempty"`
	RecoveryTemplate string `json:"recoveryTemplate,omitempty"`
	DefaultTemplate bool `json:"defaultTemplate,omitempty"`
	Type string `json:"type,omitempty"`
}

// QuerySNSTextTemplateView QuerySNSTextTemplate
type QuerySNSTextTemplateView struct {
	Inventories []SNSTextTemplateInventoryView `json:"inventories,omitempty"`
}

// CreateSNSTextTemplateEventView CreateSNSTextTemplateEvent
type CreateSNSTextTemplateEventView struct {
	Inventory SNSTextTemplateInventoryView `json:"inventory,omitempty"`
}

// UpdateSNSTextTemplateEventView UpdateSNSTextTemplateEvent
type UpdateSNSTextTemplateEventView struct {
	Inventory SNSTextTemplateInventoryView `json:"inventory,omitempty"`
}

// DeleteSNSTextTemplateEventView DeleteSNSTextTemplateEvent
type DeleteSNSTextTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

