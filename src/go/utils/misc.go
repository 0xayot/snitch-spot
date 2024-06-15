package utils

func GetString(data map[string]interface{}, key string) string {
	if value, ok := data[key]; ok {
		if strValue, ok := value.(string); ok {
			return strValue
		}
	}
	return ""
}

func GetFloat64(data map[string]interface{}, key string) float64 {
	if value, ok := data[key]; ok {
		if floatValue, ok := value.(float64); ok {
			return floatValue
		}
	}
	return 0
}
