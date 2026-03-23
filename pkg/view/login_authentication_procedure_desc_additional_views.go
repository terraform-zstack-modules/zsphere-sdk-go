// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoginAuthenticationProcedureDescView LoginAuthenticationProcedureDesc
type LoginAuthenticationProcedureDescView struct {
	Order int `json:"order,omitempty"`
	Name string `json:"name,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

