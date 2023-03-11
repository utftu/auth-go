package auth

import (
	authGoCore "auth-go/auth-go-core"
	"auth-go/auth-go-core/strategies/do"
	"auth-go/auth-go-core/strategies/google"
	"auth-go/auth-go-core/strategies/github"

)

func SelectStrategy(name string, data *authGoCore.StrategyData) *authGoCore.AuthGoCore {
	switch name {
	case "do":
		return &authGoCore.AuthGoCore{
			Data:     *data,
			Strategy: &do.Strategy {},
		}
	case "google":
		return &authGoCore.AuthGoCore{
			Data: *data,
			Strategy: &google.Strategy{},
		}
	case "github": {
		return &authGoCore.AuthGoCore{
			Data: *data,
			Strategy: &github.Strategy{},
		}
	}
	default:
		return nil
	}
}