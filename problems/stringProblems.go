package problems

// Is Unique: determine if a string has all unique characters
func IsUniqueCharString(s string) bool {
	table := make(map[rune]bool, 50)

	for _, c := range s {
		if table[c] == false {
			table[c] = true
			continue
		}
		return false
	}
	return true
}
