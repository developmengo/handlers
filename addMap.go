package handlers

import "fmt"

func Addmap(a map[string]interface{}, b map[string]interface{}) map[string]interface{} {
	for k, v := range b {
		a[k] = v
	}

	return a
}

func Parse(prefix string, m map[string]interface{}) string {
	if len(prefix) > 0 {
		prefix = prefix + "."
	}

	var builder string
	for mKey, mVal := range m {

		pp := prefix + mKey

		switch typedVal := mVal.(type) {
		case string:
			builder += fmt.Sprintf("%s%s, ", pp, typedVal)
		case float64:
			builder += fmt.Sprintf("%s.%-1.0f, ", pp, typedVal)
		case map[string]interface{}:
			builder += Parse(pp, typedVal)
		}
	}

	return builder
}
