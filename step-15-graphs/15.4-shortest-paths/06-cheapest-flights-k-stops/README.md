# Cheapest Flights Within K Stops

**Problem:** `flights[i] = [from, to, price]`. Find the cheapest price from `src` to `dst` using **at most `k` stops** (so at most `k + 1` flights). Return `-1` if impossible.

```
n = 4, flights = [0→1 100, 1→2 100, 2→0 100, 1→3 600, 2→3 200], src 0, dst 3, k = 1
0→1→3 costs 700 (1 stop) ✅    0→1→2→3 costs 400 but needs 2 stops ✗
answer: 700
```

## Why plain Dijkstra fails
Dijkstra keeps only the cheapest way into each node. But a cheaper route might use **more stops**, and then it can't continue. A pricier route with fewer stops might be the only one that can still reach `dst`. So "cheapest so far" isn't the whole story: stops matter too.

## Approach 1 (Striver): BFS by number of stops
Put `(stops, node, cost)` into a **plain queue**. Because each step adds exactly 1 stop, the queue is naturally sorted by stops (level by level), so no heap is needed.
- Keep `cost[node]` = cheapest cost found so far. Push a neighbour only if the new cost is lower.
- Don't expand anything with `stops > k`.
- **Time:** O(E · k) roughly, **Space:** O(V + E)

## Approach 2: Bellman-Ford limited to k + 1 rounds (very clean)
Round `i` of Bellman-Ford finds the cheapest prices using at most `i` edges. Run exactly `k + 1` rounds. The catch: within one round, read prices from a **copy** of the previous round, or a single round could chain several flights.
```
price = [INF...], price[src] = 0
repeat k+1 times:
    next = copy(price)
    for each flight (u, v, w):
        if price[u] != INF and price[u] + w < next[v]: next[v] = price[u] + w
    price = next
```
- **Time:** O(k · E), **Space:** O(V)

## Dry run (Bellman-Ford, k = 1 → 2 rounds)
- start: [0 ∞ ∞ ∞]
- round 1: 1 = 100 → [0 100 ∞ ∞]
- round 2 (reads round 1): 2 = 200, 3 = 700 → [0 100 200 700]
- answer `price[3] = 700` ✅ (the 400 route would need a 3rd round)
