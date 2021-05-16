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

func areAllStringsUnique(slice []string) (bool, string) {
	set := make(map[string]int)
	for _, e := range slice {
		if _, ok := set[e]; ok {
			return false, e
		} else {
			set[e] = 0
		}
	}
	return true, ""
}
