// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CandidateDecisionEntryView CandidateDecisionEntry
type CandidateDecisionEntryView struct {
	Decision string `json:"decision,omitempty"`
	DecisionMaker string `json:"decisionMaker,omitempty"`
	Reason string `json:"reason,omitempty"`
}

