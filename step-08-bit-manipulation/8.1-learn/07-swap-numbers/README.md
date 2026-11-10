# Swap Two Numbers (without a temp variable)

## XOR swap
```
a = a ^ b
b = a ^ b   // = (a^b)^b = a
a = a ^ b   // = (a^b)^a = b
```
Works because `x ^ x = 0` and `x ^ 0 = x`.

**Careful:** if `a` and `b` are the **same variable** (same memory), the first line makes it 0. Two different variables with the same value are fine.

## In real Go code
Just write `a, b = b, a`. The XOR swap is an interview trick to show you understand XOR, nothing more.
