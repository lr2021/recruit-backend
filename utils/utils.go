package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"regexp"
)

func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func Base64Encode(str string) string {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(str string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func CheckUserName(userID string) bool {
	re := regexp.MustCompile(`^[0-9a-zA-Z_]{3,20}$`)
	result := re.MatchString(userID)
	return result
}

func CheckStudentID(studentID string) bool {
	re := regexp.MustCompile(`^((2021\d{9})|(2020\d{9})|(2019\d{9})|(2018\d{9})|(2017\d{9}))$`)
	result := re.MatchString(studentID)
	return result
}

func CheckPhoneNumber(tel string) bool {
	re := regexp.MustCompile(`^((0\d{2,3}-\d{7,8})|(1[3658479]\d{9}))$`)
	result := re.MatchString(tel)
	return result
}

func CheckQQNumber(qqNumber string) bool {
	if len(qqNumber) == 0 {
		result := true
		return result
	}
	re := regexp.MustCompile(`^[1-9]\d{4,10}$`)
	result := re.MatchString(qqNumber)
	return result
}

func CheckEmail(email string) bool {
	if len(email) == 0 {
		result := true
		return result
	}
	re := regexp.MustCompile(`^([A-Za-z0-9_\-\.\\u4e00-\\u9fa5])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`)
	result := re.MatchString(email)
	return result
}