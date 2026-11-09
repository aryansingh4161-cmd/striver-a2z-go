# Book Allocation (Allocate Minimum Pages)

**Problem:** `pages[i]` = pages in book `i`. Give books to `m` students so that:
- each student gets at least one book,
- each student gets a **contiguous** block of books,
- every book is given out.

Minimise the **maximum** pages any student gets. Return −1 if `m > n`.

`[25 46 28 49 24]`, m = 4 → 71 (`[25 46] [28] [49] [24]`)

A **min of max** problem → binary search on the answer.

## Checker: `studentsNeeded(limit)`
Greedy: give books to the current student while the total stays ≤ `limit`; otherwise start a new student. (Same as [ship packages](../06-ship-packages/).)

## Search space
- Low = `max(pages)` (someone must hold the biggest book).
- High = `sum(pages)` (one student takes everything).
Bigger limit → fewer students needed. We want the smallest limit with `studentsNeeded(limit) <= m`.

Why `<= m` and not `== m`? If fewer students are needed, you can always split a block further to use exactly `m` (as long as `m <= n`).

- **Time:** O(n × log(sum − max)), **Space:** O(1)
