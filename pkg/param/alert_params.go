// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteAlertParamDetail DeleteAlert detail param
type DeleteAlertParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAlertParam DeleteAlert request param
type DeleteAlertParam struct {
	BaseParam
	Params DeleteAlertParamDetail `json:"deleteAlert"`
}
