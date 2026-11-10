# Disjoint Set Union (Union-Find)

## The problem it solves
You have `n` items that start in their own groups. You need two operations, many times, very fast:
- **Union(a, b):** merge the groups of `a` and `b`.
- **Find(a):** which group is `a` in? (so "are a and b connected?" = `Find(a) == Find(b)`)

Doing a BFS/DFS after every change would be O(V + E) per question. DSU answers in **almost O(1)**, even while the graph keeps changing ("dynamic connectivity").

## How it's stored
Each group is a small tree. `parent[x]` points up; the root is the group's "leader" (`parent[root] == root`).
```
Find(x): walk up parents until you reach the root.
Union(a, b): ra = Find(a), rb = Find(b); if ra != rb: make one root point to the other.
```

## Two tricks that make it fast
**1. Path compression (in Find):** after finding the root, make every node on the path point **directly** to it. Next time it's one hop.
```
Find(x):
    if parent[x] == x: return x
    parent[x] = Find(parent[x])     // compress
    return parent[x]
```

**2. Union by rank / size:** attach the **smaller** tree under the **bigger** one, so trees stay shallow.
- **By size:** `size[root]` = number of items in the group. Small goes under big; `size[big] += size[small]`. Bonus: you can ask "how big is my group?"
- **By rank:** `rank` ≈ height estimate. Lower rank goes under higher; if equal, pick one and add 1 to its rank.

Both are fine; size is usually more useful in problems.

## Dry run (union by size)
Unions: (1,2) (2,3) (4,5) (6,7) (5,6) → are 3 and 7 connected? no. Then (3,7) → yes.
```
(1,2): 2 → 1          size1 = 2
(2,3): Find(2)=1, 3 → 1   size1 = 3
(4,5): 5 → 4          size4 = 2
(6,7): 7 → 6          size6 = 2
(5,6): roots 4, 6 equal size → 6 → 4, size4 = 4
Find(3)=1, Find(7)=4 → different
(3,7): size1 = 3 < size4 = 4 → 1 → 4, size4 = 7
Find(3)=4 = Find(7) ✅
```

## Complexity
With both tricks, each operation is **O(α(n))**, where α (inverse Ackermann) is ≤ 4 for any n you'll ever see. Treat it as O(1).

## When to reach for DSU
"Merge groups", "are these connected", "count components while edges are added", Kruskal's MST, grouping equal things (accounts, stones).
