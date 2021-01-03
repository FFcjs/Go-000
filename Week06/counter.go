package Week06

import (
	"math"
	"sync"
	"time"
)

type Counter struct {
	MaxSize int64             //最大bucket大小
	Buckets map[int64]*bucket //每一秒对应一个bucket
	Sum float64
	Mutex   *sync.RWMutex
}

type bucket struct {
	Value float64
}

// NewCounter init a Counter
func NewCounter(size int64) *Counter {
	r := &Counter{
		MaxSize: size,
		Buckets: make(map[int64]*bucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

//getCurrentBucket get current bucket. if bucket is not existed , create a bucket and return
func (c *Counter) getCurrentBucket() *bucket {
	now := time.Now().Unix()
	b, ok := c.Buckets[now]
	if !ok {
		b = &bucket{}
		c.Buckets[now] = b
	}
	return b
}

//removeExpiredBuckets remove all expired buckets
func (c *Counter) removeExpiredBuckets() {
	now := time.Now().Unix() - c.MaxSize

	for timestamp := range c.Buckets {
		if timestamp <= now {
			c.Sum-=c.Buckets[timestamp].Value
			delete(c.Buckets, timestamp)
		}
	}
}

// Add add value to current bucket
func (c *Counter) Add(i float64) {
	if i == 0 {
		return
	}

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	b := c.getCurrentBucket()
	b.Value += i
	c.Sum +=i
	c.removeExpiredBuckets()
}

func (c *Counter) GetSum() float64{
	return c.Sum
}

func (c *Counter) GetAvg() float64{
	return c.Sum/float64(len(c.Buckets))
}

func (c *Counter) GetMax() float64{
	var max float64
	now:=time.Now().Unix()
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	for timestamp, bucket := range c.Buckets {
		if timestamp >= now-c.MaxSize {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}

	return max
}

func (c *Counter) GetMin() float64{
	min:=math.MaxFloat64
	now:=time.Now().Unix()
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	for timestamp, bucket := range c.Buckets {
		if timestamp >= now-c.MaxSize {
			if bucket.Value < min {
				min = bucket.Value
			}
		}
	}

	return min
}

