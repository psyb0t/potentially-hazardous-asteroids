package cache

import "time"

// Cache is a map of Item pointers
type Cache map[string]*Item

// Item is some cached data held in memory
type Item struct {
	ExpireTime time.Time
	Data       interface{}
}

// NewItem instantiates an Item struct with the given args attached
// to its fields and returns a pointer to it
func NewItem(expireTime time.Time, data interface{}) *Item {
	return &Item{ExpireTime: expireTime, Data: data}
}

// IsExpired tells if the cache Item is expired
func (i Item) IsExpired() bool {
	if i.ExpireTime.Before(time.Now()) {
		return true
	}

	return false
}
