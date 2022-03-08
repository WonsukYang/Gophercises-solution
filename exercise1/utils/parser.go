package utils

import (
	"fmt"
	s "strings"
)

func ParseSlice(slice []int) string {
	return s.Replace(s.Trim(fmt.Sprintf("%v", slice), "[]"), " ", ",", -1)
}
