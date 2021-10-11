package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lr2021/recruit-backend/general"
	"github.com/lr2021/recruit-backend/general/config"
	"github.com/lr2021/recruit-backend/general/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type CustomClaims struct {
	UserID string
	jwt.StandardClaims
}

func CheckValidation(code string, tel string, username string) bool {

	return true
}

func GenerateValidation(tel string, username string) (bool, error) {

	return false, nil
}

func CheckReCaptcha(token string) error {
	if token == "" {
		return errors.Forbidden("auth:001", "invalid token")
	}

	var verifyResponse struct {
		Success     bool   `json:"success"`
		ChallengeTs string `json:"challenge_ts"`
		Hostname    string `json:"hostname"`
		ErrorCodes  int    `json:"error-codes"`
	}

	captchaVerify, err := http.Post(
		"https://www.recaptcha.net/recaptcha/api/siteverify",
		"application/x-www-form-urlencoded",
		strings.NewReader("secret="+config.RECAPT_SECRET_KEY+"&response="+token))
	if err != nil {
		return errors.InternalServerError("auth:999", "remote auth server response error")
	}
	content, _ := ioutil.ReadAll(captchaVerify.Body)
	if err := json.Unmarshal(content, &verifyResponse); err != nil {
		return errors.InternalServerError("auth:001", "auth data parsing failed")
	}
	if !verifyResponse.Success {
		return errors.Forbidden("auth:001", "invalid token")
	}
	return nil
}

func GenerateToken(username string) string {
	maxAge := 60 * 60 * 12
	customClaims := &CustomClaims{
		UserID: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims).SignedString([]byte(general.GetStringEnv("JWT_SECRET", "ad1mm_03et2r")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n", token)

	return token
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(general.GetStringEnv("JWT_SECRET", "ad1mm_03et2r")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", tokenString)
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
