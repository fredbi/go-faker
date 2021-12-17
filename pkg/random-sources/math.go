package sources

import (
	"math/rand"
	"sync"
)

type mathSource struct {
	lk  sync.Mutex
	src rand.Source64
}

func newMathSource(seed int64) *mathSource {
	return &mathSource{
		src: rand.NewSource(seed),
	}
}

func (r *mathSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()

	return
}

func (r *mathSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src.Uint64()
	r.lk.Unlock()

	return
}

func (r *mathSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}
