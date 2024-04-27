package helpers

import (
	"fmt"
	"sort"
	"strings"
)

func FormatValue(v interface{}) string {
	switch value := v.(type) {
	case string:
		return value
	case int, int32, int64, float32, float64:
		return fmt.Sprintf("%v", value)
	case bool:
		return fmt.Sprintf("%t", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func MapToString(data map[string]interface{}) string {
	var builder strings.Builder

	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := data[key]
		formattedValue := FormatValue(value)
		builder.WriteString(fmt.Sprintf("%s - %s <br />", key, formattedValue))
	}

	return builder.String()
}
