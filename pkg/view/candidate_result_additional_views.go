// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CandidateResultView CandidateResult
type CandidateResultView struct {
	Candidate interface{} `json:"candidate,omitempty"`
	FinalDecision string `json:"finalDecision,omitempty"`
	Decisions []CandidateDecisionEntryView `json:"decisions,omitempty"`
}

