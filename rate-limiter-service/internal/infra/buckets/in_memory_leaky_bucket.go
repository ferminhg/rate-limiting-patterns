package buckets

import "log"

type InMemoryLeakyBucket struct {
	capacity int
	water    int
}

var instance *InMemoryLeakyBucket

func NewInMemoryLeakyBucket(capacity int) *InMemoryLeakyBucket {
	if instance == nil {
		instance = &InMemoryLeakyBucket{
			capacity: capacity,
			water:    0,
		}
		log.Printf("[🪣] Created new bucket instance with capacity: %d", capacity)
	}
	return instance
}

func (b *InMemoryLeakyBucket) IsFull() bool {
	return b.water >= b.capacity
}

func (b *InMemoryLeakyBucket) Inc() {
	if b.water < b.capacity {
		log.Printf("[🪣] Increasing water: %d → %d", b.water, b.water+1)
	}
	log.Printf("[🪣] Water status: %d/%d", b.water, b.capacity)
	b.water++
}

func (b *InMemoryLeakyBucket) Dec() {
	if b.water > 0 {
		log.Printf("[🪣] Decreasing water: %d → %d", b.water, b.water-1)
	}
	b.water--
}
