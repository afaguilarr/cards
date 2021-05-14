package deck

func containsString(slice []string, e string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, se := range slice {
		if e == se {
			return true
		}
	}
	return false
}
