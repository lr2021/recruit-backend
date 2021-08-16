package general

import "os"

func GetStringEnv(key string, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	} else {
		return value
	}
}
