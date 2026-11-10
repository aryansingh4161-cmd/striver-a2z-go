# Convert a Min-Heap to a Max-Heap

**Problem:** Rearrange a min-heap array in place into a valid max-heap.

`[3 5 9 6 8 20 10 12 18 9]` → a max-heap such as `[20 18 10 12 9 9 3 5 6 8]`

## Key insight
The fact that the input is a min-heap **doesn't help**. Just run the O(n) **build-heap** for a max-heap: sift down every non-leaf from `n/2 − 1` to 0.
- **Time:** O(n), **Space:** O(1)

Many valid max-heaps exist for the same values; tests should check the **property**, not one exact array.
