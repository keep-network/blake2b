package blake2

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestF_2b(t *testing.T) {
	for i, test := range testVectors_2b {
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

func TestF_2bX(t *testing.T) {
	for i, test := range testVectors_2bX {
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
