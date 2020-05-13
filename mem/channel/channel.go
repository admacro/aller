// https://golang.org/ref/mem#tmp_4
//
// limiting concurrency with semaphore
//
// The kth receive on a channel with capacity C happens before the k+Cth
// send from that channel completes.
// see ./channel_test.go and ./chan_test.go for examples
//
// This rule generalizes the previous rule to buffered channels. It allows
// a counting semaphore to be modeled by a buffered channel: the number of
// items in the channel corresponds to the number of active uses, the capacity
// of the channel corresponds to the maximum number of simultaneous uses,
// sending an item acquires the semaphore, and receiving an item releases
// the semaphore. This is a common idiom for limiting concurrency.
package channel

import (
	"sync"
)

type channel struct {
	capacity   int
	data       interface{}
	sendMtx    sync.Mutex
	receiveMtx sync.Mutex
}

func New() *channel {
	return &channel{0, nil, sync.Mutex{}, sync.Mutex{}}
}

func (c *channel) Send(d interface{}) {
	c.sendMtx.Lock()
	for c.capacity > 0 {
	}
	c.data = d
	c.capacity++
	c.sendMtx.Unlock()
}

func (c *channel) Receive() (data interface{}) {
	c.receiveMtx.Lock()
	for c.capacity <= 0 {
	}
	data = c.data
	c.capacity--
	c.receiveMtx.Unlock()
	return
}
