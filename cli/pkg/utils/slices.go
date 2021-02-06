package utils

func IsStringInSlice(haystack []string, needle string) bool {
	for _, s := range haystack {
		if needle == s {
			return true
		}
	}
	return false
}
