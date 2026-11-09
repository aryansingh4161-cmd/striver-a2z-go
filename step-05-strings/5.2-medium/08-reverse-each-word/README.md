# Reverse Each Word (keep word order)

**Problem (LeetCode 557):** Reverse the characters of every word, keeping the words in the same order and keeping the spaces.

`"Let's take LeetCode contest"` → `"s'teL ekat edoCteeL tsetnoc"`

Compare with [reverse words](../../5.1-basic/02-reverse-words/), which reverses the **order of words** instead.

## Approach: Two pointers per word (optimal)
Convert to `[]byte`. Find each word's start and end; reverse that range in place.
- **Time:** O(n), **Space:** O(n) for the byte copy (Go strings can't be changed in place)
