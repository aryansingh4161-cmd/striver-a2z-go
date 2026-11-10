# Minimum Window Subsequence

**Problem (LeetCode 727):** Return the shortest substring `W` of `s1` such that `s2` is a **subsequence** of `W` (in order, gaps allowed). On ties, the leftmost. `""` if none.

`s1 = "abcdebdde"`, `s2 = "bde"` → `"bcde"`

Different from [minimum window substring](../03-minimum-window-substring/): here **order matters**.

## Approach: Forward scan, then backward scan (optimal two-pointer)
1. **Forward:** from position `i`, walk `s1` matching `s2` character by character. When all of `s2` is matched at index `end`, we know a valid window **ends** at `end`.
2. **Backward:** from `end`, walk left matching `s2` from its **last** character. Where it finishes (`start`) is the **latest possible start** for that end: the tightest window.
3. Record `s1[start..end]` if shorter. Continue the forward search from `start + 1`.

- **Time:** O(n × m) worst case, **Space:** O(1)

## Dry run `"abcdebdde"`, `"bde"`
- forward from 0: b(1) d(3) e(4) → end = 4; backward: e(4) d(3) b(1) → window `"bcde"` (len 4)
- forward from 2: b(5) d(6) e(8) → end = 8; backward: e(8) d(7) b(5) → `"bdde"` (len 4, not shorter)
- answer **"bcde"**

(A DP solution in O(n × m) also exists; the two-pointer version is simpler and uses O(1) space.)
