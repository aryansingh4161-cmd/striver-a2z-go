# Rabin-Karp: Find All Occurrences

**Problem:** Return every index where pattern `p` (length m) occurs in text `s` (length n).

```
s = "abababab", p = "aba" → [0 2 4]
```

## Brute force
Compare `p` at every start: **O(n · m)**. Bad when there are many near-matches (`s = "aaaa…a"`, `p = "aa…ab"`).

## Idea: compare numbers, not strings
Treat a window of m characters as a number in base `B` (say 31), modulo a big prime `M`:
```
hash("abc") = (a·B² + b·B + c) mod M        (letters mapped to 1..26)
```
- If two hashes differ → the strings **definitely** differ.
- If they're equal → **probably** equal: check the characters to be sure (this protects against collisions).

## Rolling the hash in O(1)
Slide the window one step right: remove the leftmost character, shift, add the new one.
```
h = (h − s[i] · B^(m−1)) · B + s[i+m]      (all mod M)
```
Precompute `B^(m−1) mod M` once.

## Dry run (`s = "abab…"`, `p = "aba"`, B = 31, values a=1 b=2)
- hash(p) = 1·961 + 2·31 + 1 = 1024
- window "aba" at 0: 1024 → equal → verify chars ✓ → **0**
- roll to "bab": (1024 − 1·961)·31 + 2 = 1955 ≠ 1024
- roll to "aba" at 2: (1955 − 2·961)·31 + 1 = 1024 → verify ✓ → **2**
- ... and **4** ✅

## Complexity
- Average / expected: **O(n + m)** (collisions are rare with a big prime).
- Worst case (adversarial collisions): O(n · m). Using two different moduli ("double hashing") makes that practically impossible.

Use Go `int` (64-bit) with `M ≈ 1e9 + 7` so `h · B` never overflows.
