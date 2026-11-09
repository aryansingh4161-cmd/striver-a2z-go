# Majority Element II (> n/3 times)

**Problem (LeetCode 229):** Return all elements that appear **more than n/3** times.

`[1 1 1 3 3 2 2 2]` → `[1 2]`

**Key fact:** at most **2** elements can appear more than n/3 times (three of them would need more than n elements).

## Approach 1: Count each (brute)
O(n²).

## Approach 2: Hash map (better)
Count; collect keys with count > n/3. Stop early once you have 2. **Time O(n), Space O(n)**.

## Approach 3: Extended Moore's Voting (optimal)
Track **two** candidates with two counters:
```
for x in a:
    if count1 == 0 and x != cand2: cand1 = x; count1 = 1
    else if count2 == 0 and x != cand1: cand2 = x; count2 = 1
    else if x == cand1: count1++
    else if x == cand2: count2++
    else: count1--; count2--
```
Then **verify** both candidates with a second pass (they might not really be > n/3).
- **Time:** O(n), **Space:** O(1)

The `x != cand2` check stops both slots from holding the same value.
