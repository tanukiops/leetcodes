package easy

func KthDistinct(arr []string, k int) string {
	var distinctStrings []string
	for _, v := range arr {
		if isDistinct(v, arr) {
			distinctStrings = append(distinctStrings, v)
		}
	}
	if k > len(distinctStrings) {
		return ""
	}
	return distinctStrings[k-1]
}

func isDistinct(s string, arr []string) bool {
	count := 0
	for _, v := range arr {
		if v == s {
			count++
		}
	}
	return count == 1
}
