package utils

// pathLengths returns the lengths of each path in number of edges.
func pathLengths(paths [][]*Room) []int {
	lengths := make([]int, len(paths))
	for i, p := range paths {
		lengths[i] = len(p) - 1
	}
	return lengths
}

// initialCounts computes how many ants can start on each path based on the turn count.
func initialCounts(lengths []int, t int) []int {
	counts := make([]int, len(lengths))
	for i, l := range lengths {
		c := t - l + 1
		if c < 0 {
			c = 0
		}
		counts[i] = c
	}
	return counts
}

// adjustCounts ensures that total ants assigned does not exceed the number of ants.
func adjustCounts(counts []int, ants int) []int {
	total := 0
	for _, c := range counts {
		total += c
	}
	for i := len(counts) - 1; total > ants && i >= 0; i-- {
		if counts[i] > 0 {
			take := counts[i]
			if take > total-ants {
				take = total - ants
			}
			counts[i] -= take
			total -= take
		}
	}
	return counts
}

// orderFromCounts returns the order of path indices for ant departure.
func orderFromCounts(counts []int) []int {
	var order []int
	active := true
	n := len(counts)
	for active {
		active = false
		for i := 0; i < n; i++ {
			if counts[i] > 0 {
				order = append(order, i)
				counts[i]--
				active = true
			}
		}
	}
	return order
}
