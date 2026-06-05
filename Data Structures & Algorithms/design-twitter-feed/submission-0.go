type Twitter struct {
	time      int
	tweetMap  map[int][]Tweet
	followMap map[int]map[int]struct{}
}

type Tweet struct {
	id   int
	time int
}

func Constructor() Twitter {
	return Twitter{
		time:      0,
		tweetMap:  make(map[int][]Tweet),          // userId -> list of tweets
		followMap: make(map[int]map[int]struct{}), // userId -> set of followeeId
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	if t.tweetMap[userId] == nil {
		t.tweetMap[userId] = make([]Tweet, 0)
	}
	t.tweetMap[userId] = append(
		t.tweetMap[userId], Tweet{
			id:   tweetId,
			time: t.time,
		},
	)
	t.time++
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	result := make([]int, 0, 10)
	if t.followMap[userId] == nil {
		t.followMap[userId] = make(map[int]struct{})
	}
	t.followMap[userId][userId] = struct{}{}
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	followees := t.followMap[userId]
	for followee := range followees {
		followeeTweets := t.tweetMap[followee]
		for i := len(followeeTweets)-1; i >= 0; i-- {
			heap.Push(maxHeap, followeeTweets[i])
		}
	}
	for i := 0; i < 10 && len(*maxHeap) > 0; i++ {
		result = append(result, heap.Pop(maxHeap).(Tweet).id)
	}
	return result
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.followMap[followerId]; !ok {
		t.followMap[followerId] = make(map[int]struct{})
	}
	t.followMap[followerId][followeeId] = struct{}{}
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if t.followMap[followerId] != nil {
		delete(t.followMap[followerId], followeeId)
	}
}

type MaxHeap []Tweet

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].time > h[j].time
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Tweet))
}
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}
