package helper

import (
	"strconv"
)

func Multiple7(num int) string {
	if num%7 == 0 {
		return strconv.Itoa(num) + " adalah kelipatan 7"
	} else {
		return strconv.Itoa(num) + " bukan kelipatan 7"
	}
}