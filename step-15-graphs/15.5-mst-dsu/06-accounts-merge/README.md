# Accounts Merge

**Problem:** Each account is `[name, email1, email2, ...]`. Two accounts belong to the same person if they share **any** email. Merge them. Output each person as `[name, sorted emails...]`.

```
["John", "j1@x", "j2@x"]
["John", "j3@x"]
["John", "j2@x", "j4@x"]   ← shares j2 with the first
["Mary", "m@x"]
→ ["John", "j1@x", "j2@x", "j4@x"], ["John", "j3@x"], ["Mary", "m@x"]
```
Note: same name doesn't mean same person (the two Johns without a shared email stay separate).

## Idea: DSU over account indices
1. Walk every email. Keep `owner[email]` = first account index that had it.
2. If an email is seen again in account `i`, `Union(i, owner[email])`: they're the same person.
3. Group emails by `Find(owner[email])`.
4. For each group: sort emails, put the name in front.

## Dry run (example above)
- Account 0: j1 → 0, j2 → 0
- Account 1: j3 → 1
- Account 2: j2 already owned by 0 → Union(2, 0); j4 → 2
- Account 3: m → 3
- Groups: root 0 → {j1, j2, j4}, root 1 → {j3}, root 3 → {m} ✅

- **Time:** O(N log N) where N = total emails (sorting dominates), **Space:** O(N)

Order of the output accounts doesn't matter; the test sorts them before comparing.
