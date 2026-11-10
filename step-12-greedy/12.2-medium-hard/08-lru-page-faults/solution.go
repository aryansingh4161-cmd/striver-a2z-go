package lrupagefaults

// PageFaults counts misses with LRU replacement. Time O(n * capacity).
func PageFaults(pages []int, capacity int) int {
	mem := []int{} // least recent at index 0
	faults := 0
	for _, p := range pages {
		pos := -1
		for i, x := range mem {
			if x == p {
				pos = i
				break
			}
		}
		if pos >= 0 { // hit: move to most recent
			mem = append(mem[:pos], mem[pos+1:]...)
		} else {
			faults++
			if len(mem) == capacity {
				mem = mem[1:]
			}
		}
		mem = append(mem, p)
	}
	return faults
}
