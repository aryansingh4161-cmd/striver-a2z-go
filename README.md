# Striver's A2Z DSA Sheet in Go

My notes and solutions for **Striver's A2Z DSA Sheet** (take U forward), written in **Go**.
Every problem has simple notes (brute → better → optimal), a dry run, complexity, Go code and tests you can run.

> The explanations are written from scratch in simple words. The order follows the official sheet so each folder matches a video.

---

## How this repo is organised

```
striver-a2z-go/
├── docs/go-for-dsa.md          ← C++/Java → Go translation guide (read first)
├── step-01-basics/             ← each step has a README with the theory
│   ├── README.md
│   ├── 1.4-basic-maths/
│   │   ├── README.md           ← topic summary
│   │   └── 01-count-digits/
│   │       ├── README.md       ← problem notes
│   │       ├── solution.go     ← Go solution (all approaches)
│   │       └── solution_test.go← tests
│   └── ...
├── step-02-sorting/
└── ...
```

## How to use it

1. Watch the lecture in the [A2Z playlist](https://www.youtube.com/playlist?list=PLgUwDviBIf0oF6QL8m22w1hIDC1vJ_BHz).
2. **Try the problem yourself first** (LeetCode / GFG link is in the sheet).
3. Read the problem's `README.md` here: compare your approach, check the dry run.
4. Read `solution.go`, then run the tests:

```bash
go test ./...                                    # everything
go test ./step-03-arrays/...                     # one step
go test -v ./step-03-arrays/3.2-medium/01-two-sum/  # one problem
```

5. Change the code, break it, run the tests again. That's how it sticks.

**Need Go 1.22+** (uses the built-in `min` / `max`).

## Progress

| Step | Topic | Theory | Problems |
|---|---|---|---|
| 1 | [Learn the basics](step-01-basics/) | ✅ | ✅ 22 patterns + 17 problems |
| 2 | [Sorting](step-02-sorting/) | ✅ | ✅ 7 algorithms |
| 3 | [Arrays](step-03-arrays/) | ✅ | ✅ 39 problems (easy, medium, hard) |
| 4 | [Binary search](step-04-binary-search/) | ✅ | ⏳ |
| 5 | [Strings (basic + medium)](step-05-strings/) | ✅ | ⏳ |
| 6 | [Linked list](step-06-linked-list/) | ✅ | ⏳ |
| 7 | [Recursion patterns](step-07-recursion/) | ✅ | ⏳ |
| 8 | [Bit manipulation](step-08-bit-manipulation/) | ✅ | ⏳ |
| 9 | [Stack and queue](step-09-stack-queue/) | ✅ | ⏳ |
| 10 | [Sliding window and two pointers](step-10-sliding-window-two-pointers/) | ✅ | ⏳ |
| 11 | [Heaps](step-11-heaps/) | ✅ | ⏳ |
| 12 | [Greedy](step-12-greedy/) | ✅ | ⏳ |
| 13 | [Binary trees](step-13-binary-trees/) | ✅ | ⏳ |
| 14 | [Binary search trees](step-14-bst/) | ✅ | ⏳ |
| 15 | [Graphs](step-15-graphs/) | ✅ | ⏳ |
| 16 | [Dynamic programming](step-16-dynamic-programming/) | ✅ | ⏳ |
| 17 | [Tries](step-17-tries/) | ✅ | ⏳ |
| 18 | [Strings (advanced)](step-18-strings-advanced/) | ✅ | ⏳ |

✅ done · ⏳ coming in the next batches

## What every problem page contains

| Section | What it gives you |
|---|---|
| Problem + examples | short statement in plain words |
| Approaches | brute → better → optimal, each with time and space |
| Dry run | a table you can follow with a pen |
| Edge cases | the inputs that usually break solutions |
| `solution.go` | every approach as a separate function |
| `solution_test.go` | table-driven tests (some compare against a brute force on random inputs) |

## Study tips

- **Don't skip the brute force.** Interviewers want to hear the ladder brute → better → optimal.
- **Dry run on paper** before trusting code.
- **Write the complexity** for every solution you write.
- **Revise:** the "Key idea" column in each step's README is your one-line revision sheet.
- Rough pace: Steps 1–3 in week 1–2, then about one step every 3–4 days. Graphs and DP need the most time.

## Other resources

- Official sheet (problem links + written articles): [takeuforward.org A2Z sheet](https://takeuforward.org/dsa/strivers-a2z-sheet-learn-dsa-a-to-z)
- Full A2Z playlist: [YouTube](https://www.youtube.com/playlist?list=PLgUwDviBIf0oF6QL8m22w1hIDC1vJ_BHz)
- Topic series: [Recursion](https://www.youtube.com/playlist?list=PLgUwDviBIf0rGlzIn_7rsaR2FQ5e6ZOL9) · [Trees](https://www.youtube.com/playlist?list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk) · [Graphs](https://www.youtube.com/playlist?list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn) · [DP](https://www.youtube.com/playlist?list=PLgUwDviBIf0qUlt5H_kiKYaNSqJ81PMMY) · [Tries](https://youtube.com/playlist?list=PLgUwDviBIf0pcIDCZnxhv0LkHf5KzG9zp)

All credit for the curriculum goes to Raj Vikramaditya (Striver) and the take U forward team.
