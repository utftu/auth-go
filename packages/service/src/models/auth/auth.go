package auth

import (
	authGoCore "core"
	"core/strategies/do"
	"core/strategies/github"
	"core/strategies/google"
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