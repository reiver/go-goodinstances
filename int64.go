package goodinstances

import (
	"github.com/reiver/go-iter/int64"

	"math"
)

var (
	sliceInt64 = []int64{
		math.MinInt64,
		-1,
		0,
		1,
		math.MaxInt64,
	}
)

// Int64 returns an itertor over int64, with "good" instances of the int64 type.
//
// These could be useful when writing tests.
//
// Example
//
//	iterator := goodinstances.Int64()
//	
//	defer iterator.Close()
//	
//	for iterator.Next() {
//		var datum int64
//		
//		if err := iterator.Decode(&datum); nil != err {
//			return err
//		}
//
//		// ...
//	}
//	if err := iterator.Err(); nil != err {
//		return err
//	}
func Int64() Iterator {

	slice := append([]int64(nil), sliceInt64...)
	for i:=0; i<4; i++ {
		slice = append(slice,      randomness.Int63())
		slice = append(slice, -1 * randomness.Int63())

		slice = append(slice,      randomness.Int63() / (256))
		slice = append(slice, -1 * randomness.Int63() / (256))

		slice = append(slice,      randomness.Int63() / (256 * 256))
		slice = append(slice, -1 * randomness.Int63() / (256 * 256))

		slice = append(slice,      randomness.Int63() / (256 * 256 * 256))
		slice = append(slice, -1 * randomness.Int63() / (256 * 256 * 256))

		slice = append(slice,      randomness.Int63() / (256 * 256 * 256 * 256))
		slice = append(slice, -1 * randomness.Int63() / (256 * 256 * 256 * 256))

		slice = append(slice,      randomness.Int63() / (256 * 256 * 256 * 256 * 256))
		slice = append(slice, -1 * randomness.Int63() / (256 * 256 * 256 * 256 * 256))

		slice = append(slice,      randomness.Int63() / (256 * 256 * 256 * 256 * 256 * 256))
		slice = append(slice, -1 * randomness.Int63() / (256 * 256 * 256 * 256 * 256 * 256))

		slice = append(slice,      randomness.Int63() / (256 * 256 * 256 * 256 * 256 * 256 * 256))
		slice = append(slice, -1 * randomness.Int63() / (256 * 256 * 256 * 256 * 256 * 256 * 256))
	}

	iterator := iterint64.Slice{
		Slice: slice,
	}

	return &iterator
}
