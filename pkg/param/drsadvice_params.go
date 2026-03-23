// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ApplyDRSAdviceParamDetail ApplyDRSAdvice detail param
type ApplyDRSAdviceParamDetail struct {
}

// ApplyDRSAdviceParam ApplyDRSAdvice request param
type ApplyDRSAdviceParam struct {
	BaseParam
	Params ApplyDRSAdviceParamDetail `json:"applyDRSAdvice"`
}
