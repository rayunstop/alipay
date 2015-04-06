package utils

import (
	"net/url"
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
