// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateAliyunSmsSNSTextTemplateParamDetail CreateAliyunSmsSNSTextTemplate detail param
type CreateAliyunSmsSNSTextTemplateParamDetail struct {
	Sign string `json:"sign" validate:"required"`
	AlarmTemplateCode string `json:"alarmTemplateCode" validate:"required"`
	EventTemplateCode string `json:"eventTemplateCode" validate:"required"`
	EventTemplate *string `json:"eventTemplate,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ApplicationPlatformType string `json:"applicationPlatformType" validate:"required"`
	Subject *string `json:"subject,omitempty"`
	RecoverySubject *string `json:"recoverySubject,omitempty"`
	Template string `json:"template" validate:"required"`
	RecoveryTemplate *string `json:"recoveryTemplate,omitempty"`
	DefaultTemplate *bool `json:"defaultTemplate,omitempty"`
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunSmsSNSTextTemplateParam CreateAliyunSmsSNSTextTemplate request param
type CreateAliyunSmsSNSTextTemplateParam struct {
	BaseParam
	Params CreateAliyunSmsSNSTextTemplateParamDetail `json:"params"`
}
// UpdateAliyunSmsSNSTextTemplateParamDetail UpdateAliyunSmsSNSTextTemplate detail param
type UpdateAliyunSmsSNSTextTemplateParamDetail struct {
	AlarmTemplateCode *string `json:"alarmTemplateCode,omitempty"`
	Sign *string `json:"sign,omitempty"`
	EventTemplateCode *string `json:"eventTemplateCode,omitempty"`
	EventTemplate *string `json:"eventTemplate,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Subject *string `json:"subject,omitempty"`
	RecoverySubject *string `json:"recoverySubject,omitempty"`
	Template *string `json:"template,omitempty"`
	RecoveryTemplate *string `json:"recoveryTemplate,omitempty"`
	DefaultTemplate *bool `json:"defaultTemplate,omitempty"`
}

// UpdateAliyunSmsSNSTextTemplateParam UpdateAliyunSmsSNSTextTemplate request param
type UpdateAliyunSmsSNSTextTemplateParam struct {
	BaseParam
	Params UpdateAliyunSmsSNSTextTemplateParamDetail `json:"updateAliyunSmsSNSTextTemplate"`
}
