# Step 3: Arrays

Arrays are where DSA really starts. Almost every later topic (binary search, sliding window, DP) is built on the patterns you learn here.

## Core ideas (learn these once, use them everywhere)

### 1. Brute → Better → Optimal
Interviewers expect you to talk through this ladder:
- **Brute:** try everything (usually nested loops).
- **Better:** remove repeated work (hashing, sorting).
- **Optimal:** use a property of the problem (sorted order, only positives, a mathematical trick).

### 2. Two pointers
Two indexes moving through the array, usually from both ends or one chasing the other.
Used in: remove duplicates, move zeros, union of sorted arrays, 2Sum (sorted), 3Sum, 4Sum, sort 0-1-2.

### 3. Hashing (map / set)
"Have I seen X before?" in O(1).
Used in: 2Sum, longest consecutive sequence, majority element (better approach).

### 4. Prefix sums
`prefix[i] = a[0] + … + a[i]`. Sum of `a[l..r]` = `prefix[r] − prefix[l−1]`.
Combined with a map: "count/longest subarray with sum K" in O(n), even with negatives.
The same trick works with XOR.

### 5. Kadane's algorithm
Drop a running sum the moment it becomes negative. Max subarray sum in O(n).

### 6. Voting algorithms (Moore)
Majority element in O(1) space by cancelling pairs.

### 7. Merge-sort counting
Count pairs across two sorted halves during the merge: count inversions, reverse pairs.

### 8. In-place matrix tricks
Transpose + reverse (rotate), boundary pointers (spiral), first row/column as markers (set zeroes).

## Problems

### 3.1 Easy
| # | Problem | Key idea |
|---|---|---|
| 1 | [Largest element](3.1-easy/01-largest-element/) | current champion |
| 2 | [Second largest](3.1-easy/02-second-largest/) | two champions, one pass |
| 3 | [Check if sorted](3.1-easy/03-check-sorted/) | compare neighbours |
| 4 | [Remove duplicates from sorted](3.1-easy/04-remove-duplicates-sorted/) | two pointers |
| 5 | [Left rotate by one](3.1-easy/05-left-rotate-by-one/) | save first, shift |
| 6 | [Left rotate by D](3.1-easy/06-left-rotate-by-d/) | reversal algorithm |
| 7 | [Move zeros to end](3.1-easy/07-move-zeros-to-end/) | two pointers |
| 8 | [Linear search](3.1-easy/08-linear-search/) | scan |
| 9 | [Union / intersection of sorted arrays](3.1-easy/09-union-of-sorted-arrays/) | merge-style two pointers |
| 10 | [Missing number](3.1-easy/10-missing-number/) | sum formula / XOR |
| 11 | [Max consecutive ones](3.1-easy/11-max-consecutive-ones/) | running count |
| 12 | [Single number](3.1-easy/12-single-number/) | XOR |
| 13 | [Longest subarray sum K (positives)](3.1-easy/13-longest-subarray-sum-k-positives/) | sliding window |
| 14 | [Longest subarray sum K (any sign)](3.1-easy/14-longest-subarray-sum-k/) | prefix sum + map |

### 3.2 Medium
| # | Problem | Key idea |
|---|---|---|
| 1 | [Two sum](3.2-medium/01-two-sum/) | hash map / two pointers |
| 2 | [Sort 0s, 1s, 2s](3.2-medium/02-sort-0s-1s-2s/) | Dutch National Flag |
| 3 | [Majority element (> n/2)](3.2-medium/03-majority-element/) | Moore's voting |
| 4 | [Max subarray sum + print it](3.2-medium/04-kadane-max-subarray-sum/) | Kadane |
| 5 | [Stock buy and sell](3.2-medium/05-stock-buy-sell/) | min so far |
| 6 | [Rearrange by sign](3.2-medium/06-rearrange-by-sign/) | even/odd write pointers |
| 7 | [Next permutation](3.2-medium/07-next-permutation/) | break point, swap, reverse |
| 8 | [Leaders in an array](3.2-medium/08-leaders-in-array/) | scan from the right |
| 9 | [Longest consecutive sequence](3.2-medium/09-longest-consecutive-sequence/) | hash set, start points |
| 10 | [Set matrix zeroes](3.2-medium/10-set-matrix-zeroes/) | first row/col markers |
| 11 | [Rotate matrix 90°](3.2-medium/11-rotate-matrix/) | transpose + reverse |
| 12 | [Spiral traversal](3.2-medium/12-spiral-matrix/) | four boundaries |
| 13 | [Count subarrays with sum K](3.2-medium/13-count-subarrays-sum-k/) | prefix sum + count map |

### 3.3 Hard
| # | Problem | Key idea |
|---|---|---|
| 1 | [Pascal's triangle](3.3-hard/01-pascals-triangle/) | nCr built step by step |
| 2 | [Majority element (> n/3)](3.3-hard/02-majority-element-n3/) | Moore's with 2 candidates |
| 3 | [3Sum](3.3-hard/03-three-sum/) | sort + two pointers |
| 4 | [4Sum](3.3-hard/04-four-sum/) | sort + 2 loops + two pointers |
| 5 | [Largest subarray with 0 sum](3.3-hard/05-largest-subarray-zero-sum/) | prefix sum first index |
| 6 | [Count subarrays with XOR K](3.3-hard/06-count-subarrays-xor-k/) | prefix XOR + map |
| 7 | [Merge overlapping intervals](3.3-hard/07-merge-overlapping-intervals/) | sort by start |
| 8 | [Merge sorted arrays without extra space](3.3-hard/08-merge-sorted-arrays-no-extra-space/) | swap + sort / gap method |
| 9 | [Repeating and missing number](3.3-hard/09-repeating-and-missing/) | maths / XOR |
| 10 | [Count inversions](3.3-hard/10-count-inversions/) | merge sort counting |
| 11 | [Reverse pairs](3.3-hard/11-reverse-pairs/) | merge sort + counting pass |
| 12 | [Maximum product subarray](3.3-hard/12-maximum-product-subarray/) | prefix/suffix or max/min |
