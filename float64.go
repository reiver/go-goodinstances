package goodinstances

import (
	"github.com/reiver/go-iter/float64"

	"math"
)

var (
	sliceFloat64 = []float64{
		-1 * math.MaxFloat64,
		-1,
		-1 * math.SmallestNonzeroFloat64,
		0,
		math.SmallestNonzeroFloat64,
		1,
		math.MaxFloat64,
	}
)

// Float64 returns an itertor over float64, with "good" instances of the float64 type.
//
// These could be useful when writing tests.
//
// Example
//
//	iterator := goodinstances.Float64()
//	
//	defer iterator.Close()
//	
//	for iterator.Next() {
//		var datum float64
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
func Float64() Iterator {

	slice := append([]float64(nil), sliceFloat64...)
	for i:=0; i<10; i++ {
		slice = append(slice,      randomness.Float64())
		slice = append(slice, -1 * randomness.Float64())

		slice = append(slice,      randomness.Float64() * (math.MaxFloat64 / 64.0))
		slice = append(slice, -1 * randomness.Float64() * (math.MaxFloat64 / 64.0))

		slice = append(slice,      randomness.Float64() * (math.MaxFloat64 / 2.0))
		slice = append(slice, -1 * randomness.Float64() * (math.MaxFloat64 / 2.0))

		slice = append(slice,      randomness.Float64() * math.MaxFloat64)
		slice = append(slice, -1 * randomness.Float64() * math.MaxFloat64)
	}

	iterator := iterfloat64.Slice{
		Slice: slice,
	}

	return &iterator
}
