// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AliyunSmsSNSTextTemplateInventoryView AliyunSmsSNSTextTemplate
type AliyunSmsSNSTextTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	AlarmTemplateCode string `json:"alarmTemplateCode,omitempty"`
	Sign string `json:"sign,omitempty"`
	EventTemplateCode string `json:"eventTemplateCode,omitempty"`
	EventTemplate string `json:"eventTemplate,omitempty"`
	Description string `json:"description,omitempty"`
	ApplicationPlatformType string `json:"applicationPlatformType,omitempty"`
	Subject string `json:"subject,omitempty"`
	RecoverySubject string `json:"recoverySubject,omitempty"`
	Template string `json:"template,omitempty"`
	RecoveryTemplate string `json:"recoveryTemplate,omitempty"`
	DefaultTemplate bool `json:"defaultTemplate,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryAliyunSmsSNSTextTemplateView QueryAliyunSmsSNSTextTemplate
type QueryAliyunSmsSNSTextTemplateView struct {
	Inventories []AliyunSmsSNSTextTemplateInventoryView `json:"inventories,omitempty"`
}

// UpdateAliyunSmsSNSTextTemplateEventView UpdateAliyunSmsSNSTextTemplateEvent
type UpdateAliyunSmsSNSTextTemplateEventView struct {
	Inventory AliyunSmsSNSTextTemplateInventoryView `json:"inventory,omitempty"`
}

