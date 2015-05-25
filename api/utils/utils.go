package utils

import (
	"bytes"
	"net/url"
	"sort"
)

// KeySet 得到map的key集合
func KeySet(dict map[string]string) []string {

	s := make([]string, 0, len(dict))
	for k, _ := range dict {
		s = append(s, k)
	}
	return s
}

// BuildQuery 建立带参数URL
func BuildQuery(dict map[string]string) (val url.Values) {

	val = url.Values{}
	for k, v := range dict {
		val.Set(k, v)
	}
	return val
}

// PutAll 往map里放map
func PutAll(dest, src map[string]string) {
	for k, v := range src {
		dest[k] = v
	}
}

// PrepareContent 准备请求的报文
// 按照字典排序
func PrepareContent(dict map[string]string) string {

	s := make([]string, 0, len(dict))
	for k, _ := range dict {
		s = append(s, k)
	}
	// 排序
	sort.Strings(s)

	var buf bytes.Buffer
	for _, v := range s {
		param := dict[v]
		// 过滤掉空的key
		if param != "" {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(v + "=" + param)
		}
	}
	return buf.String()
}
