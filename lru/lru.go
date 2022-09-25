package lru

import (
	"time"
)

func GetLRUClock() int64 {
	return time.Now().UnixMilli()
}
