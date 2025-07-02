// Package utils TODO
package utils

import (
	"errors"
	"strings"
)

const (
	// Base TODO
	Base = 62
	// CharacterSet TODO
	CharacterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// Encode 转化为62进制
func Encode(num int64) string {
	if num == 0 {
		return string(CharacterSet[0])
	}
	var b []byte
	for num > 0 {
		remainder := num % Base
		num /= Base
		b = append([]byte{CharacterSet[remainder]}, b...)
	}
	return string(b)
}

// Decode 解析62进制
func Decode(s string) (int64, error) {
	if s == "" {
		return 0, errors.New("empty string")
	}

	var result int64
	for _, v := range s {
		pos := strings.IndexRune(CharacterSet, v)
		if pos == -1 {
			return 0, errors.New("invalid character: " + string(v))
		}
		result = result*Base + int64(pos)
	}
	return result, nil
}
