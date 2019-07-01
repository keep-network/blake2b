package compression

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestF(t *testing.T) {
	for i, test := range testVectorsF {
		t.Run(fmt.Sprintf("test vector %v", i), func(t *testing.T) {
			//toEthereumTestCase(test)

			h := test.hIn

			F(&h, test.m, test.c, test.f, test.rounds)

			if !reflect.DeepEqual(test.hOut, h) {
				t.Errorf("Unexpected result\nExpected: [%v]\nActual:   [%v]\n", test.hOut, h)
			}
		})
	}
}

type testVector struct {
	hIn    [8]uint64
	m      [16]uint64
	c      [2]uint64
	f      bool
	rounds uint32
	hOut   [8]uint64
}

var testVectorsF = []testVector{
	{
		hIn:    [8]uint64{0x6a09e667f2bd8948, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1, 0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3736353433323130, 0x3f3e3d3c3b3a3938, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x80, 0x0},
		f:      true,
		rounds: 1,
		hOut:   [8]uint64{0x62305ad4d48dade8, 0x269ef60a1bcb8b7b, 0xef6e479e643b5ac1, 0xf8017e6422ce89fb, 0x62d09ecaa81d095e, 0x855540dcbc07bd0f, 0xeb3d4f5e5e505412, 0x60ed930f027cfd8d},
	},
	{
		hIn:    [8]uint64{0x6a09e667f2bd8948, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1, 0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3736353433323130, 0x3f3e3d3c3b3a3938, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x80, 0x0},
		f:      false,
		rounds: 1,
		hOut:   [8]uint64{0x74aaa2510b5c66ea, 0x8a708bdea257192d, 0x9c9bc48d53c7f7e8, 0x7a51bffd24d5d9e6, 0xea0b4b2070fb3a3d, 0x211208e3f7bcb1a2, 0x4dec971cecbb62fa, 0xf1cd142745f8f4ee},
	},
	{
		hIn:    [8]uint64{0x8736a85f01b0a31d, 0xd67dcbe79b220b9b, 0x9d93e5897bb89757, 0x66b47652443b1927, 0xb715da2203c2ca1d, 0xd22b820a965ad9c, 0xf23bb0bd57a16fd5, 0x6ed7a9c5dc67af05},
		m:      [16]uint64{0x8786858483828180, 0x8f8e8d8c8b8a8988, 0x9796959493929190, 0x9f9e9d9c9b9a9998, 0xa7a6a5a4a3a2a1a0, 0xafaeadacabaaa9a8, 0xb7b6b5b4b3b2b1b0, 0xbfbebdbcbbbab9b8, 0xc7c6c5c4c3c2c1c0, 0xcfcecdcccbcac9c8, 0xd7d6d5d4d3d2d1d0, 0xdfdedddcdbdad9d8, 0xe7e6e5e4e3e2e1e0, 0xefeeedecebeae9e8, 0xf7f6f5f4f3f2f1f0, 0xfefdfcfbfaf9f8},
		c:      [2]uint64{0x17f, 0x0},
		f:      true,
		rounds: 1,
		hOut:   [8]uint64{0xf0eea67995b5e71b, 0xd02c40edf19aadd1, 0x31475949119813de, 0x4ea3072f8865bb19, 0x78037a72d2bb446d, 0x3deead8310b5fba5, 0x48088c5a99689c51, 0xc6632f660a2a2a45},
	},
	{
		hIn:    [8]uint64{0x6a09e667f2bd8948, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1, 0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3736353433323130, 0x3f3e3d3c3b3a3938, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x80, 0x0},
		f:      true,
		rounds: 5,
		hOut:   [8]uint64{0xf863686fe221e37f, 0xee57b5eb7a41266f, 0xe765d26d337fa9fe, 0x329af1ae378bae9b, 0xd61f06dd8ccceec2, 0x29349cd341b19df3, 0x99488a9ab05aec7b, 0xa3442268fd039ba4},
	},
	{
		hIn:    [8]uint64{0x6a09e667f2bd8948, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1, 0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3736353433323130, 0x3f3e3d3c3b3a3938, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x80, 0x0},
		f:      false,
		rounds: 5,
		hOut:   [8]uint64{0x3a6c5ac8d6606290, 0x708c08540cb5c425, 0xf72014c8b50c280b, 0x817a6ffb1b8b405f, 0x2001caca125182b5, 0x62b8a407a249ca6e, 0x9755b62e4324caaf, 0x4578a660d8dd4876},
	},
	{
		hIn:    [8]uint64{0x2a3c65f1dbd604cb, 0xc713d5ce8c0e5e6c, 0x3b22c76bdf73716e, 0x55525ba247a930f9, 0xe240763f1c01d22c, 0x514aa14d3f187227, 0x15405fbe7ecde36c, 0x77e565078839b7c9},
		m:      [16]uint64{0x8786858483828180, 0x8f8e8d8c8b8a8988, 0x9796959493929190, 0x9f9e9d9c9b9a9998, 0xa7a6a5a4a3a2a1a0, 0xafaeadacabaaa9a8, 0xb7b6b5b4b3b2b1b0, 0xbfbebdbcbbbab9b8, 0xc7c6c5c4c3c2c1c0, 0xcfcecdcccbcac9c8, 0xd7d6d5d4d3d2d1d0, 0xdfdedddcdbdad9d8, 0xe7e6e5e4e3e2e1e0, 0xefeeedecebeae9e8, 0xf7f6f5f4f3f2f1f0, 0xfefdfcfbfaf9f8},
		c:      [2]uint64{0x17f, 0x0},
		f:      true,
		rounds: 5,
		hOut:   [8]uint64{0x7e7acc657731d1df, 0xb4229114f545597b, 0x7fa4b3140217b470, 0xe86a8c27ab0b1ce8, 0xc103e219a3a2f77f, 0xa495b8c706b8016f, 0xacf32b1cb5d83f2a, 0xe8d3248892d5d74c},
	},
	{
		hIn:    [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
		m:      [16]uint64{0x3020100, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x84, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0x7180f3083d5aaabe, 0x569cd951d62cf43, 0x1dc9f9ff9eb4d014, 0xa5ef0eec4192b524, 0xba8b0407d49601f6, 0x48b0bc8e8246218d, 0x6d4fbb56fd42888d, 0xacb8aa4d4b9ce1f8},
	},
	{
		hIn:    [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1413121110, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x95, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0x99cb773e2cc379a, 0x8876af0e51773691, 0xd22f53d314339be8, 0x45292a02de394c76, 0xddf87a51130d71b5, 0x1cec3be7246631c0, 0x3620302852f17de6, 0xdd18d2b40cab30f3},
	},
	{
		hIn:    [8]uint64{0x6a09e667f2bd8948, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1, 0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3736353433323130, 0x3f3e3d3c3b3a3938, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x80, 0x0},
		f:      false,
		rounds: 12,
		hOut:   [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
	},
	{
		hIn:    [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1a1918, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x9b, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0x54777d6e31afa286, 0xac6453272e941b20, 0xd7d85bab6289ea12, 0xf9c8fffbc56d27fb, 0x67df67484eae8ca2, 0x2709162425b7d980, 0xb5e078605bda55c8, 0x1dcab91ce391aa54},
	},
	{
		hIn:    [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1b1a1918, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0x9c, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0x52780a0caf0bd10, 0x793f8aaf9b3606e7, 0xa77506a8030a2cd7, 0x64515ea4e30bb0bb, 0x6d6fb5ef88eed124, 0x6577e2e65a547757, 0x8930fc93e4f5a8c3, 0x55eedfa133896315},
	},
	{
		hIn:    [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3130, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		c:      [2]uint64{0xb2, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0xfe47a4b9e2a1e7d0, 0x1080ffe977223ee8, 0xaa7afa12ae75f3c2, 0x6aa2687831a6a58c, 0x32cfc1fb690b7a36, 0x1630637eb345da5, 0x75ba0e2310213d6f, 0x7cf5ac546fa52840},
	},
	{
		hIn:    [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3736353433323130, 0x3f3e3d3c3b3a3938, 0x4746454443424140, 0x4f4e4d4c4b4a4948, 0x5756555453525150, 0x5f5e5d5c5b5a5958, 0x6766656463626160, 0x6f6e6d6c6b6a6968, 0x7776757473727170, 0x7f7e7d7c7b7a7978},
		c:      [2]uint64{0x100, 0x0},
		f:      false,
		rounds: 12,
		hOut:   [8]uint64{0xb02406865b80cde9, 0xfec1e321559891be, 0x456158a6d71cf0ea, 0xc1eb845c70d4a3d8, 0x948e82cc975dcbd6, 0x984780bab7e23638, 0x2fb585df5bc4b043, 0xcb3e866544413d92},
	},
	{
		hIn:    [8]uint64{0xb02406865b80cde9, 0xfec1e321559891be, 0x456158a6d71cf0ea, 0xc1eb845c70d4a3d8, 0x948e82cc975dcbd6, 0x984780bab7e23638, 0x2fb585df5bc4b043, 0xcb3e866544413d92},
		m:      [16]uint64{0x8786858483828180, 0x8f8e8d8c8b8a8988, 0x9796959493929190, 0x9f9e9d9c9b9a9998, 0xa7a6a5a4a3a2a1a0, 0xafaeadacabaaa9a8, 0xb7b6b5b4b3b2b1b0, 0xbfbebdbcbbbab9b8, 0xc7c6c5c4c3c2c1c0, 0xcfcecdcccbcac9c8, 0xd7d6d5d4d3d2d1d0, 0xdfdedddcdbdad9d8, 0xe7e6e5e4e3e2e1e0, 0xe8, 0x0, 0x0},
		c:      [2]uint64{0x169, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0xca769859d42c76d4, 0xdb449924feb8b275, 0x36b9da1f74ce7ad2, 0xeb0f4625e4c6cb16, 0xe1838ccade7d451, 0x567f4a02897cc47f, 0xeae4fd8d87db1a19, 0xfe0e61a2f52322d6},
	},
	{
		hIn:    [8]uint64{0xb02406865b80cde9, 0xfec1e321559891be, 0x456158a6d71cf0ea, 0xc1eb845c70d4a3d8, 0x948e82cc975dcbd6, 0x984780bab7e23638, 0x2fb585df5bc4b043, 0xcb3e866544413d92},
		m:      [16]uint64{0x8786858483828180, 0x8f8e8d8c8b8a8988, 0x9796959493929190, 0x9f9e9d9c9b9a9998, 0xa7a6a5a4a3a2a1a0, 0xafaeadacabaaa9a8, 0xb7b6b5b4b3b2b1b0, 0xbfbebdbcbbbab9b8, 0xc7c6c5c4c3c2c1c0, 0xcfcecdcccbcac9c8, 0xd7d6d5d4d3d2d1d0, 0xdfdedddcdbdad9d8, 0xe7e6e5e4e3e2e1e0, 0xefeeedecebeae9e8, 0xf3f2f1f0, 0x0},
		c:      [2]uint64{0x174, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0x8187dd8f261496b3, 0xd5b489bffe2c5e51, 0x34e626c210ab2b40, 0x6c0dfb00e09a6b4e, 0xea800ec83e2fcb79, 0x168969f8d28019eb, 0x51653672749f2ebd, 0x37a823cad39c6416},
	},
	{
		hIn:    [8]uint64{0xb02406865b80cde9, 0xfec1e321559891be, 0x456158a6d71cf0ea, 0xc1eb845c70d4a3d8, 0x948e82cc975dcbd6, 0x984780bab7e23638, 0x2fb585df5bc4b043, 0xcb3e866544413d92},
		m:      [16]uint64{0x8786858483828180, 0x8f8e8d8c8b8a8988, 0x9796959493929190, 0x9f9e9d9c9b9a9998, 0xa7a6a5a4a3a2a1a0, 0xafaeadacabaaa9a8, 0xb7b6b5b4b3b2b1b0, 0xbfbebdbcbbbab9b8, 0xc7c6c5c4c3c2c1c0, 0xcfcecdcccbcac9c8, 0xd7d6d5d4d3d2d1d0, 0xdfdedddcdbdad9d8, 0xe7e6e5e4e3e2e1e0, 0xefeeedecebeae9e8, 0xf7f6f5f4f3f2f1f0, 0xf8},
		c:      [2]uint64{0x179, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0xf5b5dbc5e76321b3, 0xef75c8ead211dc1f, 0x7e0a0999767ecbbb, 0x5daf9507d5a8f87f, 0xcdf83e5498ffd974, 0x872785043dc19af8, 0x567417c800efe056, 0x37758ee39fd5e161},
	},
	{
		hIn:    [8]uint64{0xff3e48ac606d2724, 0x2b6d0f03be548f10, 0xff2d7f8bccfba4b8, 0x72deac9005e263eb, 0xe6d620e38bba2bf1, 0x33164736804d08f5, 0xdbf004f0c01f6df4, 0xbc02b6384c348cde},
		m:      [16]uint64{0x706050403020100, 0xf0e0d0c0b0a0908, 0x1716151413121110, 0x1f1e1d1c1b1a1918, 0x2726252423222120, 0x2f2e2d2c2b2a2928, 0x3736353433323130, 0x3f3e3d3c3b3a3938, 0x4746454443424140, 0x4f4e4d4c4b4a4948, 0x5756555453525150, 0x5f5e5d5c5b5a5958, 0x6766656463626160, 0x6f6e6d6c6b6a6968, 0x7776757473727170, 0x7f7e7d7c7b7a7978},
		c:      [2]uint64{0x100, 0x0},
		f:      false,
		rounds: 12,
		hOut:   [8]uint64{0xb02406865b80cde9, 0xfec1e321559891be, 0x456158a6d71cf0ea, 0xc1eb845c70d4a3d8, 0x948e82cc975dcbd6, 0x984780bab7e23638, 0x2fb585df5bc4b043, 0xcb3e866544413d92},
	},
	{
		hIn:    [8]uint64{0xb02406865b80cde9, 0xfec1e321559891be, 0x456158a6d71cf0ea, 0xc1eb845c70d4a3d8, 0x948e82cc975dcbd6, 0x984780bab7e23638, 0x2fb585df5bc4b043, 0xcb3e866544413d92},
		m:      [16]uint64{0x8786858483828180, 0x8f8e8d8c8b8a8988, 0x9796959493929190, 0x9f9e9d9c9b9a9998, 0xa7a6a5a4a3a2a1a0, 0xafaeadacabaaa9a8, 0xb7b6b5b4b3b2b1b0, 0xbfbebdbcbbbab9b8, 0xc7c6c5c4c3c2c1c0, 0xcfcecdcccbcac9c8, 0xd7d6d5d4d3d2d1d0, 0xdfdedddcdbdad9d8, 0xe7e6e5e4e3e2e1e0, 0xefeeedecebeae9e8, 0xf7f6f5f4f3f2f1f0, 0xfefdfcfbfaf9f8},
		c:      [2]uint64{0x17f, 0x0},
		f:      true,
		rounds: 12,
		hOut:   [8]uint64{0xccfc282ed6092714, 0x5b46f8d0fa97afd0, 0x7010c51d20821e97, 0x48923ea42a37a0fa, 0x609a13be7c1e14b, 0x6e10a4b63d85d1d5, 0x6d3d370d80f97b0a, 0x61a4f22ed6462dee},
	},
}

// toEthereumTestCase transforms F test vector into test vector format used by
// go-ethereum precompiles
func toEthereumTestCase(vector testVector) {
	var memory [213]byte

	// for h (512 bits = 64 bytes)
	binary.BigEndian.PutUint64(memory[0:8], vector.hIn[0])
	binary.BigEndian.PutUint64(memory[8:16], vector.hIn[1])
	binary.BigEndian.PutUint64(memory[16:24], vector.hIn[2])
	binary.BigEndian.PutUint64(memory[24:32], vector.hIn[3])

	binary.BigEndian.PutUint64(memory[32:40], vector.hIn[4])
	binary.BigEndian.PutUint64(memory[40:48], vector.hIn[5])
	binary.BigEndian.PutUint64(memory[48:56], vector.hIn[6])
	binary.BigEndian.PutUint64(memory[56:64], vector.hIn[7])

	// for m (1024 bits = 128 bytes)
	binary.BigEndian.PutUint64(memory[64:72], vector.m[0])
	binary.BigEndian.PutUint64(memory[72:80], vector.m[1])
	binary.BigEndian.PutUint64(memory[80:88], vector.m[2])
	binary.BigEndian.PutUint64(memory[88:96], vector.m[3])

	binary.BigEndian.PutUint64(memory[96:104], vector.m[4])
	binary.BigEndian.PutUint64(memory[104:112], vector.m[5])
	binary.BigEndian.PutUint64(memory[112:120], vector.m[6])
	binary.BigEndian.PutUint64(memory[120:128], vector.m[7])

	binary.BigEndian.PutUint64(memory[128:136], vector.m[8])
	binary.BigEndian.PutUint64(memory[136:144], vector.m[9])
	binary.BigEndian.PutUint64(memory[144:152], vector.m[10])
	binary.BigEndian.PutUint64(memory[152:160], vector.m[11])

	binary.BigEndian.PutUint64(memory[160:168], vector.m[12])
	binary.BigEndian.PutUint64(memory[168:176], vector.m[13])
	binary.BigEndian.PutUint64(memory[176:184], vector.m[14])
	binary.BigEndian.PutUint64(memory[184:192], vector.m[15])

	// 8 bytes for t[0], 8 bytes for t[1]
	binary.BigEndian.PutUint64(memory[192:200], vector.c[0])
	binary.BigEndian.PutUint64(memory[200:208], vector.c[1])

	// 1 byte for f
	if vector.f {
		memory[208] = 1
	}

	// 4 bytes for rounds
	binary.BigEndian.PutUint32(memory[209:213], uint32(vector.rounds))

	fmt.Printf("input: \"%v\"\n", hex.EncodeToString(memory[:]))

	var result [64]byte

	binary.BigEndian.PutUint64(result[0:8], vector.hOut[0])
	binary.BigEndian.PutUint64(result[8:16], vector.hOut[1])
	binary.BigEndian.PutUint64(result[16:24], vector.hOut[2])
	binary.BigEndian.PutUint64(result[24:32], vector.hOut[3])

	binary.BigEndian.PutUint64(result[32:40], vector.hOut[4])
	binary.BigEndian.PutUint64(result[40:48], vector.hOut[5])
	binary.BigEndian.PutUint64(result[48:56], vector.hOut[6])
	binary.BigEndian.PutUint64(result[56:64], vector.hOut[7])

	fmt.Printf("expected: \"%v\"\n", hex.EncodeToString(result[:]))
}
