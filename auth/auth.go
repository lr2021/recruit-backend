package auth

import "github.com/lr2021/recruit-backend/general/errors"
import "github.com/lr2021/recruit-backend/general/config"

func CheckReCaptcha(token string) error {
	if token == "" {
		return errors.Forbidden("auth:001", "invalid token")
	}
	return nil
}
