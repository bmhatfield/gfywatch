package overwatch

func contains(keywords []string, keyword string) bool {
	for _, word := range keywords {
		if word == keyword {
			return true
		}
	}

	return false
}
