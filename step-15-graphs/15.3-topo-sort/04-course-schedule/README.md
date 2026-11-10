# Course Schedule I and II

**Problem:** There are `n` courses. `prerequisites[i] = [a, b]` means "to take `a` you must first take `b`".
- **I:** Can you finish all courses? (true / false)
- **II:** Return one valid order to take them, or an empty list if impossible.

## Turning it into a graph
Each pair `[a, b]` is an edge **`b → a`** (b comes before a). Watch the direction: it's the most common mistake here.

Then:
- Finishing all courses is possible ⇔ the graph has **no cycle**.
- A valid order is a **topological order**.

So both parts are just [Kahn's algorithm](../02-kahns-algorithm/).

## Dry run
`n = 4`, prereqs `[[1,0],[2,0],[3,1],[3,2]]` → edges `0→1, 0→2, 1→3, 2→3`
- indeg 0:0 → order 0, then 1, 2, then 3 → `[0 1 2 3]`, length 4 = n → possible ✅

`n = 2`, `[[1,0],[0,1]]` → `0→1, 1→0`: cycle → I: false, II: `[]`

- **Time:** O(V + E), **Space:** O(V + E)

## Edge cases
- No prerequisites: any order works (we return 0..n-1).
- Duplicate pairs: fine, in-degree counts them twice and they get removed twice.
