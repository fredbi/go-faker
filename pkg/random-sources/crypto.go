package sources

import (
	crand "crypto/rand"
	"encoding/binary"
	"sync"
)

type cryptoRand struct {
	sync.Mutex
	buf []byte
}

func (c *cryptoRand) Seed(seed int64) {}

func (c *cryptoRand) Uint64() uint64 {
	// Lock to make reading thread safe
	c.Lock()
	defer c.Unlock()

	crand.Read(c.buf)
	return binary.BigEndian.Uint64(c.buf)
}

func (c *cryptoRand) Int63() int64 {
	return int64(c.Uint64() & ^uint64(1<<63))
}
