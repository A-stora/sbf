package bloom

import (
	"errors"
	"hash/fnv"
)

type hashFunc func([]byte) uint64

type sbf struct {
	b []uint64
	m uint
	k uint
	h1 hashFunc
	h2 hashFunc
}

func NewSBF(m uint, k uint, h1 hashFunc, h2 hashFunc) (*sbf, error) {
	// size of bitarray could be 63 bits higher than anticipated due to rounding to uint64 boundaries
	if m == 0 {
        return nil, errors.New("m must be greater than 0")
    }
    if k == 0 {
        return nil, errors.New("k must be greater than 0")
    }
	
	return &sbf{
		b: make([]uint64, (m+63)/64),
		m: m,
		k: k,
		h1: h1,
		h2: h2,
	}, nil
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

func (sbf *sbf) Add(data []byte) {
	hashes := sbf.getKHashes(data)

	for i := range hashes {
		// use whole bitarray even if larger than provided m
		pos := hashes[i] % uint64(len(sbf.b)*64)
		bucket := pos/64
		pos %= 64
		sbf.b[bucket] |= uint64(1) << pos
	}
}


func (sbf *sbf) Contains(data []byte) bool {
	// false return means definitely not in the set
	// true means probably not in the set
	
	hashes := sbf.getKHashes(data)
	for i := range hashes {
		pos := hashes[i] % uint64(len(sbf.b)*64)
		bucket := pos/64
		pos %= 64

		mask := uint64(1) << pos
		isSet := (sbf.b[bucket] & mask) != 0
		if !isSet {
			return false
		}
	}

	return true // go should add a ~maybe type
}

func (sbf *sbf) initialHashPair(data []byte) (uint64, uint64) {
    return sbf.h1(data), sbf.h2(data)
}

func (sbf *sbf) getKHashes(data []byte) []uint64 {
    hash1, hash2 := sbf.initialHashPair(data) 
    hashes := make([]uint64, sbf.k)
    
    for i := 0; i < int(sbf.k); i++ {
        hashes[i] = hash1 + uint64(i)*hash2
    }
    return hashes
}