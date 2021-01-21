package utils

import (
	"fmt"
	"strconv"
)

// StrToInt for one-liner code
func StrToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		LogToFile(fmt.Sprint(err))
	}
	return result
}

// IntToStr for one-liner code
func IntToStr(num int) string {
	result := strconv.Itoa(num)
	return result
}
