package utils

import (
	"errors"
	"strings"
)

// ParseCookies parse cookie string to cookie options
// func ParseCookies(cookie string, v interface{}) (err error) {
// 	if cookie == "" {
// 		return errors.New("cookie is empty")
// 	}
// 	list := strings.Split(cookie, "; ")
// 	cookieM := make(map[string]string, len(list))
// 	for _, v := range list {
// 		item := strings.Split(v, "=")
// 		cookieM[item[0]] = item[1]
// 	}
// 	err = mapstructure.Decode(cookieM, v)
// 	return
// }

// ParseCookiesMap parse cookie string to map
func ParseCookiesMap(cookie string) (cookies map[string]string, err error) {
	if cookie == "" {
		return nil, errors.New("cookie is empty")
	}
	list := strings.Split(cookie, "; ")
	cookieM := make(map[string]string, len(list))
	for _, v := range list {
		item := strings.Split(v, "=")
		cookieM[item[0]] = item[1]
	}
	cookies = cookieM
	return
}
