package utils

// Duplicate 去重
func Duplicate[T comparable](list []T) []T {
	m := make(map[T]bool)
	result := make([]T, 0, len(list))
	for _, t := range list {
		if !m[t] {
			result = append(result, t)
			m[t] = true
		}
	}
	return result
}
