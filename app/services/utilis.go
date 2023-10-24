package services

func CheckEmptyString(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
} 