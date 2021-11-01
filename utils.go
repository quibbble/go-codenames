package go_codenames

func contains(items []string, item string) bool {
	for _, it := range items {
		if it == item {
			return true
		}
	}
	return false
}

func indexOf(items []string, item string) int {
	for index, it := range items {
		if it == item {
			return index
		}
	}
	return -1
}
