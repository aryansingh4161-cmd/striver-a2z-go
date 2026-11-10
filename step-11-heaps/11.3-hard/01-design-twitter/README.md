# Design Twitter

**Problem (LeetCode 355):**
- `PostTweet(userId, tweetId)`
- `GetNewsFeed(userId)`: the 10 most recent tweet IDs from the user and everyone they follow, newest first
- `Follow(a, b)` / `Unfollow(a, b)`

## Design
- A global **timestamp** counter: every tweet gets `time++`.
- `tweets[user]`: slice of `(time, tweetId)` in posting order (newest at the end).
- `following[user]`: a set.

## GetNewsFeed: merge K sorted lists with a heap
Each followed user's tweets (plus your own) are already sorted by time. This is [merge K sorted lists](../../11.2-medium/04-merge-k-sorted-lists/), taking only the first 10:
1. Push each user's **newest** tweet into a **max-heap** by time (remember user + index).
2. Pop the newest; add its id; push that user's next older tweet.
3. Stop after 10.

- **PostTweet / Follow / Unfollow:** O(1)
- **GetNewsFeed:** O(F log F) where F = number of followees (plus at most 10 more pops)

A user can't follow themselves, and unfollowing someone you don't follow does nothing.
