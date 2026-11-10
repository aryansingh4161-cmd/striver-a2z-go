# Topological Sort (DFS)

**Problem:** Given a **DAG** (directed graph with no cycles), list the nodes so that for every edge `u → v`, `u` comes **before** `v`.

Think of it as "order of tasks": `u → v` means "do u before v". Getting dressed: socks → shoes, shirt → tie.

```
5 → 0 ← 4
↓       ↓
2 → 3 → 1
```
One valid order: `5 4 2 3 1 0` (many orders can be valid).

Only DAGs have a topological order. If there's a cycle (`a → b → a`), nobody can go first.

## Idea: "finish time" with DFS
When `dfs(u)` **finishes**, every node reachable from `u` has already finished. So if we push `u` onto a stack *after* exploring all its neighbours, then everything that must come after `u` is already below it in the stack.

Pop the stack (or reverse the finish list) → topological order.

```
for every unvisited node s: dfs(s)
dfs(u):
    visited[u] = true
    for v in adj[u]: if !visited[v]: dfs(v)
    stack.push(u)        // after all children
answer = stack popped top to bottom
```

## Dry run (graph above)
- dfs(0): no edges → push 0. Stack: [0]
- dfs(1): push 1. Stack: [0 1]
- dfs(2): → 3 → 1 (visited). push 3, push 2. Stack: [0 1 3 2]
- dfs(4): 0, 1 visited → push 4. Stack: [0 1 3 2 4]
- dfs(5): 2, 0 visited → push 5. Stack: [0 1 3 2 4 5]
- Pop from top: **5 4 2 3 1 0** ✅

- **Time:** O(V + E), **Space:** O(V)

## Edge cases
- Isolated nodes still appear in the order.
- If the graph might have a cycle, use Kahn's algorithm (next problem) — it can tell you.

## How to test any topo order
Don't compare with one fixed answer. Instead check: every node appears once, and for every edge `u → v`, `pos[u] < pos[v]`. The test file does exactly this.
