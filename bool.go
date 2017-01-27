package goodinstances

import (
	"github.com/reiver/go-iter/bool"
)

var (
	sliceBool = []bool{
		true,
		false,
	}
)

// Bool returns an itertor over bool, with "good" instances of the bool type.
//
// These could be useful when writing tests.
//
// Example
//
//	iterator := goodinstances.Bool()
//	
//	defer iterator.Close()
//	
//	for iterator.Next() {
//		var datum bool
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
func Bool() Iterator {

	slice := append([]bool(nil), sliceBool...)

	iterator := iterbool.Slice{
		Slice: slice,
	}

	return &iterator
}
