package config

import "github.com/lr2021/recruit-backend/general"

var (
	RECAPT_SECRET_KEY = general.GetStringEnv("RECAPT_SECRET_KEY", "")
	MYSQL_USERNAME = general.GetStringEnv("MYSQL_USERNAME", "")
	MYSQL_PASSWORD = general.GetStringEnv("MYSQL_PASSWORD", "")
	MYSQL_DBNAME = general.GetStringEnv("MYSQL_DBNAME", "")
	REDIS_PASSWORD = general.GetStringEnv("REDIS_PASSWORD", "")

	MYSQL_ADDR = general.GetStringEnv("MYSQL_ADDR", "")
	REDIS_ADDR = general.GetStringEnv("REDIS_ADDR", "")

	SMTP_HOST = general.GetStringEnv("SMTP_HOST", "")
	SMTP_PORT = general.GetStringEnv("SMTP_PORT", "")
	SMTP_USER = general.GetStringEnv("SMTP_USER", "")
	SMTP_PASS = general.GetStringEnv("SMTP_PASS", "")
)

