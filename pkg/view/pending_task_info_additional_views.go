// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PendingTaskInfoView PendingTaskInfo
type PendingTaskInfoView struct {
	Name string `json:"name,omitempty"`
	ClassName string `json:"className,omitempty"`
	Index int `json:"index,omitempty"`
	PendingTime int64 `json:"pendingTime,omitempty"`
	ExecutionTime int64 `json:"executionTime,omitempty"`
	Context string `json:"context,omitempty"`
	ApiId string `json:"apiId,omitempty"`
	ApiName string `json:"apiName,omitempty"`
	ContextList []string `json:"contextList,omitempty"`
}

