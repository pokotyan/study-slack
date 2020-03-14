package env

import (
	"context"

	settingRepository "github.com/pokotyan/study-slack/repository/setting"
)

type (
	ConnpassEnvUsecase interface {
		SetEnv(ctx context.Context, rawBody string) []EnvError
	}
	connpassEnvUsecaseImpl struct {
		settingRepo settingRepository.SettingRepository
	}
)
