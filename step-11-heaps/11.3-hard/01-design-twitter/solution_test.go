package designtwitter

import (
	"slices"
	"testing"
)

func TestTwitter(t *testing.T) {
	tw := Constructor()
	tw.PostTweet(1, 5)
	if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{5}) {
		t.Errorf("feed = %v", got)
	}
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{6, 5}) {
		t.Errorf("feed = %v", got)
	}
	tw.Unfollow(1, 2)
	if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{5}) {
		t.Errorf("feed = %v", got)
	}
	for i := 100; i < 115; i++ {
		tw.PostTweet(1, i)
	}
	if got := tw.GetNewsFeed(1); len(got) != 10 || got[0] != 114 || got[9] != 105 {
		t.Errorf("feed = %v", got)
	}
}
