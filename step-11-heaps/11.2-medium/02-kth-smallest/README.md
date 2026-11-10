# Kth Smallest Element

Mirror of [kth largest](../01-kth-largest/):
- **Max-heap of size k**: push each element; if size > k, pop the largest. The top is the k-th smallest. **O(n log k)**.
- Quickselect with target index `k − 1`. **Average O(n)**.

`[7 10 4 3 20 15]`, k = 3 → 7

Go has no built-in max-heap: flip `Less`, or push **negated** values into a min-heap (shown here, since it's a common trick you'll see in other people's code).
