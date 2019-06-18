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
			mHex, err := hex.DecodeString(test.mIn)
			if err != nil {
				t.Fatal(err)
			}

			h := test.hIn
			c := test.c

			F(&h, mHex, &c, test.f, test.rounds)

			if !reflect.DeepEqual(test.hOut, h) {
				t.Errorf("Unexpected result\nExpected: [%v]\nActual:   [%v]\n", test.hOut, h)
			}
		})
	}
}

type testVector struct {
	mIn    string
	hIn    [8]uint64
	c      [2]uint64
	f      bool
	rounds int
	hOut   [8]uint64
}
