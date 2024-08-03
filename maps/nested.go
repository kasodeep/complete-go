package maps

func GetNameCounts(names []string) map[rune]map[string]int {
	// using the nested map.
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		if name == "" {
			continue
		}

		firstChar := rune(name[0])
		_, ok := counts[firstChar]

		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}
