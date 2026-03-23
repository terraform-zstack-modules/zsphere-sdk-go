// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ElaborationContentView ElaborationContent
type ElaborationContentView struct {
	Category string `json:"category,omitempty"`
	Code string `json:"code,omitempty"`
	Regex string `json:"regex,omitempty"`
	Message_cn string `json:"message_cn,omitempty"`
	Message_en string `json:"message_en,omitempty"`
	Source string `json:"source,omitempty"`
	Method string `json:"method,omitempty"`
	Distance float64 `json:"distance,omitempty"`
}

