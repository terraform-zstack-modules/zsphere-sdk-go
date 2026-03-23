// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SmartDataStructView SmartDataStruct
type SmartDataStructView struct {
	Id int `json:"id,omitempty"`
	AttributeName string `json:"attributeName,omitempty"`
	Flag string `json:"flag,omitempty"`
	Value int `json:"value,omitempty"`
	Worst int `json:"worst,omitempty"`
	Thresh int `json:"thresh,omitempty"`
	Type string `json:"type,omitempty"`
	Updated string `json:"updated,omitempty"`
	WhenFailed string `json:"whenFailed,omitempty"`
	RawValue int64 `json:"rawValue,omitempty"`
	State string `json:"state,omitempty"`
}

