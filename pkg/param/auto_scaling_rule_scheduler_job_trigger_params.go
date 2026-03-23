// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateAutoScalingRuleSchedulerJobTriggerParamDetail CreateAutoScalingRuleSchedulerJobTrigger detail param
type CreateAutoScalingRuleSchedulerJobTriggerParamDetail struct {
	TriggerType *string `json:"triggerType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingRuleSchedulerJobTriggerParam CreateAutoScalingRuleSchedulerJobTrigger request param
type CreateAutoScalingRuleSchedulerJobTriggerParam struct {
	BaseParam
	Params CreateAutoScalingRuleSchedulerJobTriggerParamDetail `json:"params"`
}
