package utils

import (
	"strconv"
	"strings"
)

func StrinToLower(data string) string {
	return strings.ToLower(data)
}

func StringSliceToLower(slice []string) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = strings.ToLower(v)
	}
	return result
}

func MapKeysToLower(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		result[strings.ToLower(k)] = v
	}
	return result
}

func Int64ToString(item int64) string {
	return strconv.FormatInt(item, 10)
}

func ConvertInt64ToString(item int64) string {
	return strconv.FormatInt(item, 10)
}
