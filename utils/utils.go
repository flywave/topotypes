package utils

import "strings"

func StrEquals(t1, t2 string) bool {
	return strings.ToLower(t1) == strings.ToLower(t2)
}
