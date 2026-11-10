# Word Ladder II

**Problem (LeetCode 126):** Same as [Word Ladder I](../09-word-ladder/), but return **all** shortest transformation sequences.

`hit → cog` with the same list → `[hit hot dot dog cog]` and `[hit hot lot log cog]`

## Approach: BFS for levels, then DFS backwards
Storing full paths in the BFS queue explodes in memory. Instead:
1. **BFS** from `beginWord`, recording `depth[word]` (its level). Remove words from the set **level by level** (only after the whole level is processed), so several parents on the same level can reach the same word.
2. **DFS backwards** from `endWord`: move to a neighbour whose depth is exactly **one less**, until you reach `beginWord`. Each such path (reversed) is a shortest sequence.

Going backwards only follows edges that lie on shortest paths, so no wasted work.
- **Time:** O(N × L × 26) for BFS + the size of the output, **Space:** O(N × L)
