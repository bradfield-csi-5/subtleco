// Author: Patch Neranartkomol

package counterservice

import (
	"sync"
	"sync/atomic"
)

type CounterService interface {
	// Returns values in ascending order; it should be safe to call
	// getNext() concurrently from multiple goroutines without any
	// additional synchronization on the caller's side.
	getNext() uint64
}

type UnsynchronizedCounterService struct {
	/* Please implement this struct and its getNext method */
	count uint64
}

// getNext() - This one can be UNSAFE
func (counter *UnsynchronizedCounterService) getNext() uint64 {
	counter.count++
	return counter.count
}

type AtomicCounterService struct {
	/* Please implement this struct and its getNext method */
	count atomic.Uint64
}

// getNext() with sync/atomic
func (counter *AtomicCounterService) getNext() uint64 {
	return counter.count.Add(1)
}

type MutexCounterService struct {
	/* Please implement this struct and its getNext method */
	count uint64
	lock  sync.Mutex
}

// getNext() with sync/Mutex
func (counter *MutexCounterService) getNext() uint64 {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.count++
	local := counter.count
	return local
}

type ChannelCounterService struct {
	/* Please implement this struct and its getNext method */
	reqChan chan struct{}
	resChan chan uint64
}

// A constructor for ChannelCounterService
func newChannelCounterService() *ChannelCounterService {
	cs := &ChannelCounterService{
		reqChan: make(chan struct{}),
		resChan: make(chan uint64),
	}
	go cs.counterGoroutine()
	return cs
}

func (cs *ChannelCounterService) counterGoroutine() {
	var count uint64
	for range cs.reqChan {
		count++
		cs.resChan <- count
	}
}

// getNext() with goroutines and channels
func (counter *ChannelCounterService) getNext() uint64 {
	counter.reqChan <- struct{}{}
	return <-counter.resChan
}
