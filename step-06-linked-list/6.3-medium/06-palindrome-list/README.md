# Check if a Linked List is a Palindrome

**Problem (LeetCode 234):** `1 → 2 → 2 → 1` → true, `1 → 2` → false.

## Approach 1: Stack (brute)
Push all values; walk again comparing with pops (pops come out reversed). **Time O(2n), Space O(n)**.

## Approach 2: Reverse the second half (optimal)
1. Find the **first** middle with slow/fast (`for fast.Next != nil && fast.Next.Next != nil`).
2. Reverse the list after the middle.
3. Compare the first half with the reversed second half.
4. Reverse the second half back (good practice: don't leave the input broken).

- **Time:** O(n), **Space:** O(1)

## Dry run `1 → 2 → 3 → 2 → 1`
Middle = 3. Second half `2 → 1` reversed → `1 → 2`. Compare `1,2` with `1,2` ✓ → **true**.
