package helpers

func IsBetween(target, min, max int) bool {
	if (target < min) || (target > max) {
		return false
	}
	return true
}
