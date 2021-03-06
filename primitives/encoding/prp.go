// A PRP is a special form of encoding that pseudo-randomly permutes the input space.
// TODO: Should small-space PRPs be added / used instead of Shuffle?
package encoding

import (
	"crypto/rand"
	"io"
	"math/big"
)

type Shuffle struct {
	EncKey, DecKey [16]byte
}

func (s Shuffle) Encode(i byte) byte {
	return s.EncKey[i]
}

func (s Shuffle) Decode(i byte) byte {
	return s.DecKey[i]
}

func GenerateShuffle(reader io.Reader) (s Shuffle) {
	s.EncKey = [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.DecKey = [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for i := int64(15); i > 0; i-- { // Performance bottleneck.
		j, _ := rand.Int(reader, big.NewInt(i+1))
		s.EncKey[i], s.EncKey[j.Int64()] = s.EncKey[j.Int64()], s.EncKey[i]
	}

	for i, j := range s.EncKey {
		s.DecKey[j] = byte(i)
	}

	return
}
