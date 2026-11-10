package jobsequencing

import "sort"

// Job has an id, a deadline (1-based) and a profit.
type Job struct{ ID, Deadline, Profit int }

// Schedule returns the number of jobs done and the total profit.
// Time O(n log n + n * maxDeadline).
func Schedule(jobs []Job) (count, profit int) {
	js := append([]Job{}, jobs...)
	sort.Slice(js, func(i, j int) bool { return js[i].Profit > js[j].Profit })
	maxD := 0
	for _, j := range js {
		maxD = max(maxD, j.Deadline)
	}
	slot := make([]bool, maxD+1)
	for _, j := range js {
		for t := j.Deadline; t >= 1; t-- { // latest free slot
			if !slot[t] {
				slot[t] = true
				count++
				profit += j.Profit
				break
			}
		}
	}
	return count, profit
}
