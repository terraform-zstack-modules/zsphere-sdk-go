// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteBillingParamDetail DeleteBilling detail param
type DeleteBillingParamDetail struct {
	AccountUuid *string `json:"accountUuid,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBillingParam DeleteBilling request param
type DeleteBillingParam struct {
	BaseParam
	Params DeleteBillingParamDetail `json:"deleteBilling"`
}
