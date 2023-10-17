package qqhelper

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	res := RandInt(100, 999)
	t.Log(res)
	if res < 100 || res >= 999 {
		t.Errorf("[100,999)的数字，结果：%d", res)
	}
}

func TestRandString(t *testing.T) {
	length := 10
	randString := RandString(length)
	t.Log(randString)
	if len(randString) != length {
		t.Errorf("TestRandString length %d, but got %d", length, len(randString))
	}
}

func TestMd5(t *testing.T) {
	str := "123456"
	res := Md5(str)
	t.Log(res)
	if len(res) == 0 {
		t.Errorf("md5 error %s", res)
	}
}
