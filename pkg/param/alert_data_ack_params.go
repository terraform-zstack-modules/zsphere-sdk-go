// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateAlertDataAckParamDetail UpdateAlertDataAck detail param
type UpdateAlertDataAckParamDetail struct {
	ResumeAlert *bool `json:"resumeAlert,omitempty"`
}

// UpdateAlertDataAckParam UpdateAlertDataAck request param
type UpdateAlertDataAckParam struct {
	BaseParam
	Params UpdateAlertDataAckParamDetail `json:"updateAlertDataAck"`
}
