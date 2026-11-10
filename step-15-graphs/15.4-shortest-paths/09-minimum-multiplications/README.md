# Minimum Multiplications to Reach End

**Problem:** Start with number `start`. In one step pick any `a` from `arr` and do `start = (start * a) % 100000`. Fewest steps to reach `end`? `-1` if impossible.

```
arr = [2 5 7], start = 3, end = 30
3 → 3·2 = 6 → 6·5 = 30   → 2 steps
```

## Seeing the graph
- **Nodes:** the numbers `0 .. 99999` (everything is taken `% 100000`).
- **Edges:** from `x` to `(x · a) % 100000` for every `a` in `arr`, each costing 1 step.

All steps cost 1 → **BFS** gives the fewest steps. Striver writes it as "Dijkstra with a queue", which is the same thing here since every edge weight is 1.

```
dist[start] = 0; queue = [start]
pop x; for a in arr:
    y = x*a % 100000
    if dist[y] unset: dist[y] = dist[x] + 1; if y == end return; push y
```

## Edge cases
- `start == end` → 0 (check before anything).
- At most 100000 states, so BFS always ends.

- **Time:** O(100000 · len(arr)), **Space:** O(100000)
