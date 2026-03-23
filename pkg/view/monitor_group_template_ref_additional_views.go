// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MonitorGroupTemplateRefVOView MonitorGroupTemplateRefVO
type MonitorGroupTemplateRefVOView struct {
	Uuid string `json:"uuid,omitempty"`
	TemplateUuid string `json:"templateUuid,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	IsApplied bool `json:"isApplied,omitempty"`
}

