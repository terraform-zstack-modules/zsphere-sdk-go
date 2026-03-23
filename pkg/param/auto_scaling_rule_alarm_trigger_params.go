// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateAutoScalingRuleAlarmTriggerParamDetail CreateAutoScalingRuleAlarmTrigger detail param
type CreateAutoScalingRuleAlarmTriggerParamDetail struct {
	TriggerType *string `json:"triggerType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingRuleAlarmTriggerParam CreateAutoScalingRuleAlarmTrigger request param
type CreateAutoScalingRuleAlarmTriggerParam struct {
	BaseParam
	Params CreateAutoScalingRuleAlarmTriggerParamDetail `json:"params"`
}
