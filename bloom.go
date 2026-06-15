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

func NewDefaultSBF(m uint, k uint) *sbf {
	// Using standard lib FNV-1 and FNV-1a, non-cryptographic hash functions
	return &sbf{
		b: make([]uint64, (m+63)/64),
		m: m,
		k: k,
		h1: func(data []byte) uint64 {
			h := fnv.New64()
        	h.Write(data)
        	return h.Sum64()
		},
		h2: func(data []byte) uint64 {
        	h := fnv.New64a()
        	h.Write(data)
        	return h.Sum64()
    	},
	}
}

func (sbf *sbf) Add([]byte) {
	// hash the input k times

	// set correct bits in b for the hashed input
	// bucketValue |= uint64(1) << 27 
}

func (sbf *sbf) contains(data []byte) bool {
	// false return means definitely not in the set
	// true means probably not in the set
	
	// implement lookup

	return false
}

func (sbf *sbf) initialHashPair(data []byte) (uint64, uint64) {
    return sbf.h1(data), sbf.h2(data)
}

func (sbf *sbf) getKHashes(data []byte, k uint) []uint64 {
    hash1, hash2 := sbf.initialHashPair(data) 
    hashes := make([]uint64, k)
    
    for i := 0; i < int(k); i++ {
        hashes[i] = hash1 + uint64(i)*hash2
    }
    return hashes
}