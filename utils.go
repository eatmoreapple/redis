package redis

func MapToInterface(cmd string, args Q) []interface{} {
	var result = make([]interface{}, 0, len(args)+1)
	result = append(result, cmd)
	for k, v := range args {
		result = append(result, k)
		result = append(result, v)
	}
	return result
}

func StringToInterface(texts ...string) []interface{} {
	var result = make([]interface{}, 0, len(texts))
	for _, text := range texts {
		result = append(result, text)
	}
	return result
}
