// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateEventSubscriptionLabelParamDetail UpdateEventSubscriptionLabel detail param
type UpdateEventSubscriptionLabelParamDetail struct {
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
}

// UpdateEventSubscriptionLabelParam UpdateEventSubscriptionLabel request param
type UpdateEventSubscriptionLabelParam struct {
	BaseParam
	Params UpdateEventSubscriptionLabelParamDetail `json:"updateEventSubscriptionLabel"`
}
