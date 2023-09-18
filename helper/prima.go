package helper

import (
	"strconv"
)

func IsPrime(num int) string {
	output := true
	if num <= 1 {
		output = false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			output = false
		}
	}

	if !output {
		return strconv.Itoa(num) + " adalah bilangan prima"
	} else {
		return strconv.Itoa(num) + " bukan bilangan prima"
	}
}
