// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AuditsInventoryView Audits
type AuditsInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	CreateTime int64 `json:"createTime,omitempty"`
	ApiName string `json:"apiName,omitempty"`
	ClientBrowser string `json:"clientBrowser,omitempty"`
	ClientIp string `json:"clientIp,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	Operator string `json:"operator,omitempty"`
	RequestDump string `json:"requestDump,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	RequestUuid string `json:"requestUuid,omitempty"`
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
	ResponseDump string `json:"responseDump,omitempty"`
	Success bool `json:"success,omitempty"`
	SignedText string `json:"signedText,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
}

// QueryAuditView QueryAudit
type QueryAuditView struct {
	Inventories []AuditsInventoryView `json:"inventories,omitempty"`
}

