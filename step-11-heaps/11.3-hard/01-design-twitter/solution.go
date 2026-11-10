package designtwitter

import "container/heap"

type tweet struct{ time, id int }

// Twitter is a tiny social network with a merged news feed.
type Twitter struct {
	time      int
	tweets    map[int][]tweet
	following map[int]map[int]bool
}

// Constructor creates an empty Twitter.
func Constructor() Twitter {
	return Twitter{tweets: map[int][]tweet{}, following: map[int]map[int]bool{}}
}

// PostTweet records a tweet. O(1).
func (t *Twitter) PostTweet(userID, tweetID int) {
	t.time++
	t.tweets[userID] = append(t.tweets[userID], tweet{t.time, tweetID})
}

// Follow makes follower follow followee. O(1).
func (t *Twitter) Follow(follower, followee int) {
	if follower == followee {
		return
	}
	if t.following[follower] == nil {
		t.following[follower] = map[int]bool{}
	}
	t.following[follower][followee] = true
}

// Unfollow removes the follow relation. O(1).
func (t *Twitter) Unfollow(follower, followee int) {
	delete(t.following[follower], followee)
}

// cursor points at one user's tweet list position.
type cursor struct{ time, user, idx int }

type maxHeap []cursor

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(cursor)) }
func (h *maxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// GetNewsFeed returns up to 10 most recent tweet IDs (newest first).
func (t *Twitter) GetNewsFeed(userID int) []int {
	h := &maxHeap{}
	users := []int{userID}
	for f := range t.following[userID] {
		users = append(users, f)
	}
	for _, u := range users {
		if ts := t.tweets[u]; len(ts) > 0 {
			last := len(ts) - 1
			heap.Push(h, cursor{ts[last].time, u, last})
		}
	}
	feed := []int{}
	for h.Len() > 0 && len(feed) < 10 {
		c := heap.Pop(h).(cursor)
		feed = append(feed, t.tweets[c.user][c.idx].id)
		if c.idx > 0 {
			prev := t.tweets[c.user][c.idx-1]
			heap.Push(h, cursor{prev.time, c.user, c.idx - 1})
		}
	}
	return feed
}
