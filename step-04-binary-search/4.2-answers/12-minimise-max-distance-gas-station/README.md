# Minimise Maximum Distance to Gas Station

**Problem:** Gas stations at sorted positions `stations`. Add `k` new stations **anywhere** (not only at integers). Minimise the largest gap between adjacent stations. Return it (answers within 10⁻⁶ are accepted).

`[1 2 3 4 5 6 7 8 9 10]`, k = 9 → 0.5

## Approach 1: Greedy, one station at a time
Each time, put a station into the gap whose **current** section length (`gap / (placed + 1)`) is largest. With a max-heap: **O(n log n + k log n)**. Fine for small k.

## Approach 2: Binary search on a real-valued answer (optimal)
Guess the max distance `d`. For each gap `g`, how many new stations does it need so no piece exceeds `d`?
`need = ceil(g / d) − 1`
(For `g = 4, d = 1.5`: ceil(2.67) − 1 = 2 stations → pieces of 1.33.)
Total ≤ k → `d` works → try smaller.

On real numbers there's no `mid ± 1`. Instead, loop until `hi − lo < 1e-6`, or simply run a fixed **100 iterations** (each halves the range: way more precision than needed).

- **Time:** O(n × log(range / precision)), **Space:** O(1)

## Dry run idea
Gaps all 1, k = 9: one station per gap halves each gap → max 0.5. With `d = 0.5`, each gap needs `ceil(1/0.5) − 1 = 1` → total 9 ✅. With `d = 0.49`, each gap needs 2 → 18 > 9 ❌.
