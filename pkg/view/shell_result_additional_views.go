// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ShellResultView ShellResult
type ShellResultView struct {
	ReturnCode int `json:"returnCode,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	Stderr string `json:"stderr,omitempty"`
	ErrorCode ErrorCodeView `json:"errorCode,omitempty"`
}

