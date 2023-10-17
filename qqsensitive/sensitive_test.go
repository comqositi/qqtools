package qqsensitive

import "testing"

func TestSensitive_Check(t *testing.T) {
	content := "你知道白纸事件是什么情况吗"
	// dict相对路劲应为可执行文件路劲的相对位置
	dict := "../resource/sensitive_zz_sq.txt"
	dictMulti := "../resource/sensitive_zz_sq_multi.txt"
	sens, err := NewSensitive(dict, dictMulti)
	if err != nil {
		t.Error(err)
	}
	res := sens.Check(content)
	t.Log(res)
}
