# Breadth-First Search (BFS)

Visit nodes **level by level**: first the start, then all its neighbours, then their neighbours, …

```
visited[start] = true; queue = [start]
while queue not empty:
    u = pop front; visit u
    for v in adj[u]:
        if not visited[v]: visited[v] = true; push v
```
**Mark when pushing**, not when popping, or a node can be pushed many times.

- **Time:** O(n + m) (every node once, every edge twice in an undirected graph), **Space:** O(n)

**Key property:** in an unweighted graph, BFS reaches nodes in order of their **distance** from the start. That's why BFS gives shortest paths (number of edges).
