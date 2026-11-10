package stockspan

// StockSpanner answers span queries online with a monotonic stack.
type StockSpanner struct{ st [][2]int } // (price, span)

// Next records today's price and returns its span. Amortised O(1).
func (s *StockSpanner) Next(price int) int {
	span := 1
	for len(s.st) > 0 && s.st[len(s.st)-1][0] <= price {
		span += s.st[len(s.st)-1][1]
		s.st = s.st[:len(s.st)-1]
	}
	s.st = append(s.st, [2]int{price, span})
	return span
}
