// Package maximumxor has the binary trie used for XOR problems.
package maximumxor

// Bits is how many bits we store (numbers below 2^31).
const Bits = 31

type node struct{ child [2]*node }

// BitTrie stores numbers as paths of bits from the top bit down.
type BitTrie struct{ root *node }

// NewBitTrie returns an empty binary trie.
func NewBitTrie() *BitTrie { return &BitTrie{root: &node{}} }

// Insert adds x. O(31).
func (t *BitTrie) Insert(x int) {
	cur := t.root
	for b := Bits - 1; b >= 0; b-- {
		bit := x >> b & 1
		if cur.child[bit] == nil {
			cur.child[bit] = &node{}
		}
		cur = cur.child[bit]
	}
}

// MaxXor returns the largest x XOR y over stored y (the trie must not be empty). O(31).
func (t *BitTrie) MaxXor(x int) int {
	cur, val := t.root, 0
	for b := Bits - 1; b >= 0; b-- {
		bit := x >> b & 1
		if cur.child[1-bit] != nil {
			val |= 1 << b
			cur = cur.child[1-bit]
		} else {
			cur = cur.child[bit]
		}
	}
	return val
}

// Brute tries every pair. O(n²).
func Brute(nums []int) int {
	best := 0
	for i := range nums {
		for j := i; j < len(nums); j++ {
			best = max(best, nums[i]^nums[j])
		}
	}
	return best
}

// FindMaximumXOR uses the bit trie. O(31·n).
func FindMaximumXOR(nums []int) int {
	t := NewBitTrie()
	for _, x := range nums {
		t.Insert(x)
	}
	best := 0
	for _, x := range nums {
		best = max(best, t.MaxXor(x))
	}
	return best
}
