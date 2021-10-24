package go_codenames

func contains(items []string, item string) bool {
	for _, it := range items {
		if it == item {
			return true
		}
	}
	return false
}
