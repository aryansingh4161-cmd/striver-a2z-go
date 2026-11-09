# Majority Element (> n/2 times)

**Problem (LeetCode 169):** Return the element that appears **more than n/2** times. (LeetCode guarantees it exists; in general, verify.)

`[2 2 1 1 1 2 2]` → 2

## Approach 1: Count each (brute)
Nested loop counting. **Time O(n²)**.

## Approach 2: Hash map (better)
Count with a map; return the key with count > n/2. **Time O(n), Space O(n)**.

## Approach 3: Moore's Voting Algorithm (optimal)
Think of it as a battle: the majority element has more "soldiers" than all the others combined, so it survives when every soldier cancels one enemy.
```
candidate, count = ?, 0
for x in a:
    if count == 0: candidate = x; count = 1
    else if x == candidate: count++
    else: count--
```
Then (if not guaranteed) **verify** by counting `candidate` once more.
- **Time:** O(n), **Space:** O(1)

## Dry run `[7 7 5 7 5 1 5 7 5 5 7 7 5 5 5 5]`
The count drops to 0 a few times, and the candidate changes. At the end the candidate is **5**, which appears 9 times out of 16 (> 8). ✓
