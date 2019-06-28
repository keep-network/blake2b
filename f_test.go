package blake2

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestF_2b(t *testing.T) {
	runTest(t, testVectors2b)
}

func TestF_2bX(t *testing.T) {
	runTest(t, testVectors2bX)
}

func runTest(t *testing.T, testVectors []testVector) {
	for i, test := range testVectors {
		t.Run(fmt.Sprintf("test vector %v", i), func(t *testing.T) {
			mDecoded, err := hex.DecodeString(test.mIn)
			if err != nil {
				t.Fatal(err)
			}
			if len(mDecoded) != BlockSize {
				t.Fatalf("unexpected message block vector length [%v]", len(mDecoded))
			}

			var m [BlockSize]byte
			copy(m[:], mDecoded)
			h := test.hIn

			F(&h, m, test.t, test.f, test.rounds)

			if !reflect.DeepEqual(test.hOut, h) {
				t.Errorf("Unexpected result\nExpected: [%v]\nActual:   [%v]\n", test.hOut, h)
			}
		})
	}
}

type testVector struct {
	mIn    string
	hIn    [8]uint64
	t      [2]uint64
	f      bool
	rounds uint32
	hOut   [8]uint64
}
