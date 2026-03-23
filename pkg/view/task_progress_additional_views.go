// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TaskProgressInventoryView TaskProgress
type TaskProgressInventoryView struct {
	BaseInfoView
	BaseTimeView
	ApiId string `json:"apiId,omitempty"`
	Content string `json:"content,omitempty"`
	Opaque map[string]interface{} `json:"opaque,omitempty"`
	CreateTime int64 `json:"createTime,omitempty"`
	LastOpTime int64 `json:"lastOpTime,omitempty"`
	CurrentStep int64 `json:"currentStep,omitempty"`
	TotalStep int64 `json:"totalStep,omitempty"`
	TaskUuid string `json:"taskUuid,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Time int64 `json:"time,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

