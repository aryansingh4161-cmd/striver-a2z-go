# Alien Dictionary

**Problem:** You get a list of words from an alien language, **already sorted** in that language's alphabet order. Using the first `k` letters (`a`, `b`, ...), find an order of letters consistent with the list.

```
words = ["baa", "abcd", "abca", "cab", "cad"], k = 4
answer: "bdac"
```

## Step 1: Pull out letter rules from neighbouring words
Compare each word with the **next** one. Look for the **first position where they differ**: that gives one rule.

| pair | first difference | rule |
|---|---|---|
| baa / abcd | b vs a | b before a |
| abcd / abca | d vs a | d before a |
| abca / cab | a vs c | a before c |
| cab / cad | b vs d | b before d |

Only the **first** difference matters. After that, the letters tell us nothing (like "car" < "cat" says nothing about the letters after r/t).

Why only neighbours? If w1 < w2 < w3, the rule between w1 and w3 already follows from the other two.

## Step 2: Topological sort of the letters
Each rule `x before y` is an edge `x → y`. A topological order of the letters is a valid alphabet. Use [Kahn's algorithm](../02-kahns-algorithm/).

Edges: b→a, d→a, a→c, b→d → order **b d a c** ✅

## Edge cases (important!)
- **Prefix problem:** `["abc", "ab"]` is impossible: a longer word can't come before its own prefix. Return "".
- **Cycle** in the rules (e.g. a before b and b before a): no valid order, return "".
- Letters that never appear in any rule still must be in the answer (they can go anywhere).

- **Time:** O(total characters + k), **Space:** O(k + rules)
