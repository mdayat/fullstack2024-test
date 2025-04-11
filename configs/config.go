package configs

import (
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type Configs struct {
	Env      Env
	Db       Db
	Redis    *redis.Client
	Validate *validator.Validate
}

func NewConfigs(env Env, db Db, redis *redis.Client) Configs {
	return Configs{
		Env:      env,
		Db:       db,
		Redis:    redis,
		Validate: NewValidate(),
	}
}
