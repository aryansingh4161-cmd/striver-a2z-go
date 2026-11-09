# Isomorphic Strings

**Problem (LeetCode 205):** `s` and `t` are isomorphic if you can replace characters in `s` to get `t`, where:
- every occurrence of a character maps to the same character,
- two different characters can't map to the same character.

| s, t | answer |
|---|---|
| "egg", "add" | true (e→a, g→d) |
| "foo", "bar" | false (o→a and o→r) |
| "badc", "baba" | false (b→b and d→b) |

## Approach 1: Two maps
`s→t` map and `t→s` map. For each position check both directions are consistent.
- **Time:** O(n), **Space:** O(1) (at most 256 entries)

## Approach 2: "Last seen position" arrays (neat trick)
Two arrays of size 256 storing the last index (+1) where each character appeared. At each `i`, the last-seen positions of `s[i]` and `t[i]` must be **equal**; then update both to `i + 1`.
Equal positions mean both characters have the same "history", so the mapping is consistent both ways.
- **Time:** O(n), **Space:** O(1)
