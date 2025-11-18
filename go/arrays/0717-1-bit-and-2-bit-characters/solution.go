package bitand2bitcharacters

func oneBitCharacter(bits []int) bool {

	n := len(bits)
	i := 0
	for i < n-1 {
		if bits[i] == 0 {
			i += 1
		} else {
			i += 2
		}
	}
	return i == n-1
}

func isOneBitCharacter(bits []int) bool {
	return oneBitCharacter(bits)
}
