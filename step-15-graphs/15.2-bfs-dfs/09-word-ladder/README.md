# Word Ladder I

**Problem (LeetCode 127):** Transform `beginWord` into `endWord` changing **one letter at a time**; every intermediate word must be in `wordList`. Return the number of words in the shortest sequence (including both ends), or 0.

`hit → hot → dot → dog → cog` → 5

## Approach: BFS where words are nodes
Two words are neighbours if they differ by one letter. BFS gives the shortest sequence.
Generating neighbours: for each position, try all 26 letters, and keep the ones in the word **set**. Remove a word from the set as soon as it's queued (that's the "visited" mark).
- **Time:** O(N × L × 26) (N words of length L), **Space:** O(N × L)

Comparing every pair of words would be O(N² × L): much slower for big lists.
