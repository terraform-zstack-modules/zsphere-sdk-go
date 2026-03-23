// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ActionStructView ActionStruct
type ActionStructView struct {
	ResourceName string `json:"resourceName,omitempty"`
	ActionName string `json:"actionName,omitempty"`
	Round int `json:"round,omitempty"`
	InDegree []string `json:"inDegree,omitempty"`
	Actions interface{} `json:"actions,omitempty"`
}

