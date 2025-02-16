package godsc

func Compare(left, right *string) float64 {
	// One of the strings is nil
	if left == nil || right == nil {
		return 0
	}

	// Both are pointers to the same string
	if left == right {
		return 1
	}

	runesLeft := []rune(*left)
	runesRight := []rune(*right)

	// Both strings are empty
	if len(runesLeft) == 0 && len(runesRight) == 0 {
		return 1
	}

	// One of the strings is empty
	if len(runesLeft) == 0 || len(runesRight) == 0 {
		return 0
	}

	// Split the strings into bigrams
	bigramsLeft := bigrams(runesLeft)
	bigramsRight := bigrams(runesRight)

	// Calculate the intersection of the bigrams
	intersectionCount := 0
	for _, bigramLeft := range bigramsLeft {
		for _, bigramRight := range bigramsRight {
			if bigramLeft == bigramRight {
				intersectionCount++
			}
		}
	}

	// Calculate the DSC
	return 2 * float64(intersectionCount) / float64(len(bigramsLeft)+len(bigramsRight))
}

func bigrams(runes []rune) []string {
	bigrams := make([]string, len(runes)-1)
	for i := 0; i < len(runes)-1; i++ {
		bigrams[i] = string(runes[i : i+2])
	}

	return bigrams
}
