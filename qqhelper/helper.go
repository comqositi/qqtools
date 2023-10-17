package qqhelper

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// 常用函数
const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// RandInt 随机生成范围内的数字
func RandInt(start, end int) int {
	// 无需自动播种, go语言本身已播种
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return start + rand.Intn(end-start+1)
}

// RandString 随机生成字符串
func RandString(length int) string {
	b := make([]byte, length)
	for i, _ := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Md5 字符串md5
func Md5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// StringDateToTime 日期字符串转秒时间戳
func StringDateToTime(localTime string) (int64, error) {

	TimeLocation, _ := time.LoadLocation("Asia/Shanghai")
	// 解析时间字符串为time.Time对象
	t, err := time.ParseInLocation(time.DateOnly, localTime, TimeLocation)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// StringDateTimeToTime 日期时间字符串转秒时间戳
func StringDateTimeToTime(localTime string) (int64, error) {

	TimeLocation, _ := time.LoadLocation("Asia/Shanghai")
	// 解析时间字符串为time.Time对象
	t, err := time.ParseInLocation(time.DateTime, localTime, TimeLocation)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
