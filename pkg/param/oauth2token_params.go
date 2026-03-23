// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// GetOAuth2TokenParamDetail GetOAuth2Token detail param
type GetOAuth2TokenParamDetail struct {
}

// GetOAuth2TokenParam GetOAuth2Token request param
type GetOAuth2TokenParam struct {
	BaseParam
	Params GetOAuth2TokenParamDetail `json:"getOAuth2Token"`
}
