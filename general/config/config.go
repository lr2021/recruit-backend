package config

import "github.com/lr2021/recruit-backend/general"

var (
	RECAPT_SECRET_KEY = general.GetStringEnv("RECAPT_SECRET_KEY", "")
)

