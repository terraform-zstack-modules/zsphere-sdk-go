// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AuditDataView AuditData
type AuditDataView struct {
	Id int64 `json:"id,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ClientIp string `json:"clientIp,omitempty"`
	ClientBrowser string `json:"clientBrowser,omitempty"`
	ApiName string `json:"apiName,omitempty"`
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	RequestUuid string `json:"requestUuid,omitempty"`
	ResponseUuid string `json:"responseUuid,omitempty"`
	SessionUuid string `json:"sessionUuid,omitempty"`
	RequestDump string `json:"requestDump,omitempty"`
	ResponseDump string `json:"responseDump,omitempty"`
	Operator string `json:"operator,omitempty"`
	Time int64 `json:"time,omitempty"`
}

