package auth

import (
	authGoCore "auth-go/auth-go-core"
	"auth-go/auth-go-core/strategies/do"
)

func SelectStrategy(name string, data *authGoCore.StrategyData) *authGoCore.AuthGoCore {
	switch name {
	case "do":
		return &authGoCore.AuthGoCore{
			Data:     *data,
			Strategy: &do.Strategy {},
		}
	default:
		return nil
	}
}