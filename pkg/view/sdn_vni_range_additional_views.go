// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SdnVniRangeView SdnVniRange
type SdnVniRangeView struct {
	StartVni int `json:"startVni,omitempty"`
	EndVni int `json:"endVni,omitempty"`
}

