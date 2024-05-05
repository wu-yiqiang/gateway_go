package utils

func UniqueArray(arr []string) []string {
	result := []string{}
	seen := map[string]bool{}
	for _, value := range arr {
		if _, exists := seen[value]; !exists {
			result = append(result, value)
			seen[value] = true
		}
	}
	return result
}
