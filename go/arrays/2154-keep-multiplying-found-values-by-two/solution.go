package keepmultiplyingfoundvaluesbytwo

func findFinalValue(nums []int, original int) int {
	seen := make(map[int]bool)
	for _, v := range nums {
		seen[v] = true
	}

	for seen[original] {
		original *= 2
	}
	return original
}
