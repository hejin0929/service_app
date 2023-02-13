package utils

import (
	"strconv"
	"strings"
)

var RegRuler = "^1[345789]{1}\\d{9}$"

func ZHToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\u`, `u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
