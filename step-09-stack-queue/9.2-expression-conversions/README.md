# 9.2 Prefix, Infix and Postfix Conversions

All six conversions are in [`conversions.go`](conversions.go). Operands are single letters or digits; operators are `+ − * / ^`.

| Notation | Example | Operator goes |
|---|---|---|
| Infix | `(a+b)*c` | between operands (needs brackets and precedence) |
| Postfix (Reverse Polish) | `ab+c*` | after its operands |
| Prefix (Polish) | `*+abc` | before its operands |

Postfix and prefix need **no brackets**: the order alone fixes the meaning. That's why computers (and calculators) like them.

**Precedence:** `^` (3) > `* /` (2) > `+ −` (1). `^` is **right-associative**: `a^b^c = a^(b^c)`. The others are left-associative.

---

## 1. Infix → Postfix (Shunting-yard)
Scan left to right with an **operator stack**:
- operand → output
- `(` → push
- `)` → pop to output until `(`; discard `(`
- operator `op` → while the top is an operator with **higher** precedence, or **equal** precedence and `op` is left-associative, pop it to output. Then push `op`.
- at the end, pop everything to output.

`a+b*(c^d-e)` → `abcd^e-*+`
- **Time:** O(n), **Space:** O(n)

## 2. Infix → Prefix
1. **Reverse** the infix string and swap `(` ↔ `)`.
2. Run infix → postfix (with one tweak: for equal precedence, pop only if the operator is `^`, because reversing flips associativity).
3. **Reverse** the result.

`(a+b)*c` → `*+abc`

## 3. Postfix → Infix
Scan left to right with a **string stack**:
- operand → push it
- operator → pop `b`, pop `a`, push `"(" + a + op + b + ")"`

`ab+c*` → `((a+b)*c)`

## 4. Prefix → Infix
Same as 3, but scan **right to left**, and the **first** pop is the left operand:
operator → pop `a`, pop `b`, push `"(" + a + op + b + ")"`

`*+abc` → `((a+b)*c)`

## 5. Postfix → Prefix
Left to right: operator → pop `b`, pop `a`, push `op + a + b`.
`ab+c*` → `*+abc`

## 6. Prefix → Postfix
Right to left: operator → pop `a`, pop `b`, push `a + b + op`.
`*+abc` → `ab+c*`

All conversions: **Time O(n)** stack operations. (String concatenation makes them O(n²) in the worst case; fine for these sizes.)
