package utils

import "strings"

func StrEquals(t1, t2 string) bool {
	return strings.EqualFold(t1, t2)
}
