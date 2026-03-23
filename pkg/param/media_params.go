// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteMediaParamDetail DeleteMedia detail param
type DeleteMediaParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteMediaParam DeleteMedia request param
type DeleteMediaParam struct {
	BaseParam
	Params DeleteMediaParamDetail `json:"deleteMedia"`
}
