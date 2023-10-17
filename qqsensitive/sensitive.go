package qqsensitive

import (
	"errors"
	"sort"
	"strings"
)

type Sensitive struct {
	// 单模词过滤器
	filter *Filter
	// 多模词过滤器
	filterMulti *Filter
}

// NewSensitive 创建实例
func NewSensitive(dict, dictMulti string) (*Sensitive, error) {
	if dict == "" || dictMulti == "" {
		// 使用默认字典
		return nil, errors.New("未传入字典文件路劲， 字典格式每行一个词，参考：https://github.com/comqositi/qqtools/resource下的词库文件")
	}

	ses := &Sensitive{}

	ses.filter = New()
	err := ses.filter.LoadWordDict(dict)
	if err != nil {
		return nil, err
	}
	ses.filterMulti = New()
	err = ses.filterMulti.LoadWordDict(dictMulti)
	if err != nil {
		return nil, err
	}
	return ses, nil
}

// Check 默认使用单模词和多模词匹配
func (s *Sensitive) Check(content string) string {
	sens := s.filter.FindAll(content)
	if len(sens) == 0 {
		return ""
	}
	res := combineList(sens)
	for i := 0; i < len(res); i++ {
		//ok, str := s.filterMulti.Validate(res[i])
		//if !ok {
		//	return str
		//}
		ok, str := s.filterMulti.FindIn(res[i])
		if ok {
			return str
		}
	}
	return ""
}

// CombineList 字符串切片组合
func combineList(arr []string) (res []string) {

	allCombinations := make([][]string, 15, 15)

	for i := 1; i <= len(arr); i++ {
		combs := combinations(arr, i)
		for _, c := range combs {
			perms := permute(c)
			for _, p := range perms {
				allCombinations = append(allCombinations, p)
			}
		}
	}

	// 对组合按长度进行排序
	sort.Slice(allCombinations, func(i, j int) bool {
		return len(allCombinations[i]) < len(allCombinations[j])
	})

	for _, comb := range allCombinations {
		if len(comb) == 0 {
			continue
		}
		res = append(res, strings.Join(comb, "=="))
	}
	return res
}

// combinations generates all combinations of the given data.
func combinations(data []string, length int) [][]string {
	if length == 0 {
		return [][]string{{}}
	}
	if len(data) == 0 {
		return nil
	}

	first := data[0]
	rest := data[1:]

	// With first and without first
	withFirst := combinations(rest, length-1)
	for i := range withFirst {
		withFirst[i] = append([]string{first}, withFirst[i]...)
	}

	withoutFirst := combinations(rest, length)

	return append(withFirst, withoutFirst...)
}

// permute generates all permutations of the given data.
func permute(data []string) [][]string {
	if len(data) == 0 {
		return [][]string{{}}
	}

	var result [][]string
	for i, v := range data {
		// Copy current data slice and remove current element
		tmp := append([]string(nil), data...)
		tmp = append(tmp[:i], tmp[i+1:]...)

		// Recursively generate permutations for the rest of the elements
		for _, perm := range permute(tmp) {
			result = append(result, append([]string{v}, perm...))
		}
	}
	return result
}
