package bloom

import "hash/fnv"

type hashFunc func([]byte) uint64

type sbf struct {
	b []uint64
	m uint
	k uint
	h1 hashFunc
	h2 hashFunc
}

func NewSBF(m uint, k uint, h1 hashFunc, h2 hashFunc) *sbf {
	// size of bitarray could be 63 bits higher than anticipated due to rounding to uint64 boundaries
	return &sbf{
		b: make([]uint64, (m+63)/64),
		m: m,
		k: k,
		h1: h1,
		h2: h2,
	}
}

func (sbf *sbf) Add([]byte) {
	// hash the input k times

	// set correct bits in b for the hashed input
}

func (sbf *sbf) contains(data []byte) bool {
	// false return means definitely not in the set
	// true means probably not in the set
	
	// implement lookup
	
	return false
}

func hash(data []byte) uint64 {
	h := fnv.New64()
	h.Write(data)
	return h.Sum64()
}
