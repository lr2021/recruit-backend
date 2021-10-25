package auth

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-gomail/gomail"
	"github.com/go-redis/redis/v7"
	"github.com/lr2021/recruit-backend/general"
	"github.com/lr2021/recruit-backend/general/config"
	"github.com/lr2021/recruit-backend/general/db/cache"
	"github.com/lr2021/recruit-backend/general/errors"
	"github.com/lr2021/recruit-backend/general/mail"
	"github.com/lr2021/recruit-backend/utils"
	"log"
	"math/rand"
	"strconv"

	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type CustomClaims struct {
	UserID string
	jwt.StandardClaims
}

func CheckValidation(code string, goal string, username string) bool {
	result, err := cache.GetRDB().Get(goal).Result()
	if err != nil || result != code {
		return false
	}
	return true
}

func GenerateValidation(goal string, username string) (bool, error) {
	if ok := utils.CheckEmail(goal); ok {
		code := generateCode()
		content := fmt.Sprintf("你好%v，欢迎注册凌睿工作室，你的验证码是%v,请在5分钟内完成注册，请勿泄露验证码！", username, code)
		result, err := cache.GetRDB().Set(goal, code, 300 * time.Second).Result()
		if err != nil {
			return false, err
		}
		if result == "OK" {
			res, err := sendEmail(goal, content)
			return res, err
		}
		return false, nil
	}
	if ok := utils.CheckPhoneNumber(goal); ok {
		code := generateCode()
		content := fmt.Sprintf("你好%v，欢迎注册凌睿工作室，你的验证码是%v,请在5分钟内完成注册，请勿泄露验证码！", username, code)
		result, err := cache.GetRDB().Set(goal, code, 300 * time.Second).Result()
		if err != nil {
			return false, err
		}
		if result == "OK" {
			res, err := sendLetter(goal, content)
			return res, err
		}
		return false, nil
	}
	return false, fmt.Errorf("invalid email or phone number")
}

func sendEmail(goal string, content string) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "recruit@lingruistudio.com")
	m.SetHeader("To", goal)
	m.SetHeader("Subject", "凌睿工作室招新平台注册验证邮件")
	m.SetBody("text/html", content)

	return mail.Send(m)
}

func sendLetter(goal string, code string) (bool, error) {
	client, err := dysmsapi.NewClientWithAccessKey(config.ALIYUN_REGION, config.ALIYUN_ACCESS_KEY_ID, config.ALIYUN_ACCESS_KEY_SECRET)
	if client == nil {
		return false, err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = goal
	request.SignName = "凌睿工作室"
	request.TemplateCode = "SMS_195863875"
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	response, err := client.SendSms(request)
	if err != nil {
		return false, err
	}

	log.Println(response.Code)
	if response.Code != "OK" {
		return false, fmt.Errorf("%v", response.Code)
	}
	return true, nil
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code)
	return res
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
	cache.GetRDB().Set(username, token, time.Duration(maxAge) * time.Second)
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

func RemoveToken(username string) error {
	token, err := cache.GetRDB().Get(username).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.Forbidden("auth:002", "cannot find token")
		} else {
			return err
		}
	}
	claims, err := ParseToken(token)
	if err != nil {
		return err
	}
	if claims.UserID == username {
		result, err := cache.GetRDB().Del(username).Result()
		if err != nil {
			return err
		}
		if result == 1 {
			return nil
		}
		return errors.InternalServerError("auth:999", "remote auth server response error")
	}
	return errors.Forbidden("auth:001", "invalid token")
}
