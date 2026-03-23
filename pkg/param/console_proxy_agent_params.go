// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ReconnectConsoleProxyAgentParamDetail ReconnectConsoleProxyAgent detail param
type ReconnectConsoleProxyAgentParamDetail struct {
	AgentUuids []string `json:"agentUuids,omitempty"`
}

// ReconnectConsoleProxyAgentParam ReconnectConsoleProxyAgent request param
type ReconnectConsoleProxyAgentParam struct {
	BaseParam
	Params ReconnectConsoleProxyAgentParamDetail `json:"reconnectConsoleProxyAgent"`
}
// UpdateConsoleProxyAgentParamDetail UpdateConsoleProxyAgent detail param
type UpdateConsoleProxyAgentParamDetail struct {
	ConsoleProxyOverriddenIp string `json:"consoleProxyOverriddenIp" validate:"required"`
	ConsoleProxyPort *int `json:"consoleProxyPort,omitempty"`
}

// UpdateConsoleProxyAgentParam UpdateConsoleProxyAgent request param
type UpdateConsoleProxyAgentParam struct {
	BaseParam
	Params UpdateConsoleProxyAgentParamDetail `json:"updateConsoleProxyAgent"`
}
