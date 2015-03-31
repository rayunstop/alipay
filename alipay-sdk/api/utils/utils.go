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
func BuildQuery(serverURL string, dict map[string]string) string {

	u, err := url.Parse(serverURL)
	if err != nil {
		return ""
	}
	q := u.Query()
	for k, v := range dict {
		fmt.Println(k)
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String()
}
