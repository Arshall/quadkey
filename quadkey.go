// Generates Quadkeys (xxxx-xxxx-xxxx-.......)
// Setup: go get github.com/onef9day/quadkey

package quadkey

import (
	"bytes"
	"fmt"
	"math/rand"
)

const (
	dash                 byte = '-'
	defaultAlphaNumerics      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits             = 6                    // 6 bits to represent a letter index
	letterIdxMask             = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax              = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// QuadKey contains data for generating Quadkeys
type QuadKey struct {
	RandSrc  rand.Source
	Alphabet string
	Segments int
}

func (qk *QuadKey) random(n int) []byte {
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, qk.RandSrc.Int63(), letterIdxMax; i >= 0; {

		if remain == 0 {
			cache, remain = qk.RandSrc.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(qk.Alphabet) {
			b[i] = qk.Alphabet[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

// Init creates a new QuadKey data
func Init(rndSrc rand.Source) *QuadKey {
	return &QuadKey{
		RandSrc:  rndSrc,
		Alphabet: defaultAlphaNumerics,
		Segments: 4,
	}
}

// Key Generates QuadKey in the format xxxx-xxxx-xxxx-xxxx
// eg. QuadKey(4) will generate QuadKey with 4 segments xxxx-xxxx-xxxx-xxxx
// where xxxx represents alphanumeric string
func (qk *QuadKey) Key() string {
	var b bytes.Buffer

	if qk.RandSrc == nil {
		panic(fmt.Errorf("\"QuadKey.RandSrc\" can't be nil"))
	}

	if qk.Segments <= 0 {
		panic(fmt.Errorf("can't generate Quadkey with zero segments"))
	} else if qk.Segments == 1 {
		return string(qk.random(4))
	}

	for l := 1; l <= (2*qk.Segments - 1); l++ {

		if l%2 != 0 {
			b.Write(qk.random(4))

		} else {
			b.WriteByte(dash)
		}
	}
	return b.String()
}
