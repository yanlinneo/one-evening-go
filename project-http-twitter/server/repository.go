package server

import "sync"

type TweetsRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
	rwLock sync.RWMutex
}

func (t *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	t.rwLock.RLock()
	defer t.rwLock.RUnlock()

	return t.tweets, nil
}

func (t *TweetsMemoryRepository) AddTweet(tw Tweet) (int, error) {
	t.rwLock.Lock()
	defer t.rwLock.Unlock()

	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}
