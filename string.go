package goodinstances

import (
	"github.com/reiver/go-iter/string"

	"fmt"
)

var (
	sliceString = []string{
		"",
		" ",
		"\r\n",
		"Hello world!",
		"apple",
		"banana",
		"cherry",
		"date",
		"appleBananaCherryDate",
		"APPLE_BANANA_CHERRY_DATE",
		"Apple-Banana-Cherry-Date",
		"apple banana cherry date",
		"AppleBananaCherryDate",
		"apple_banana_cherry_date",
		"Apple Banana Cherry Date",
		"APPLE BANANA CHERRY DATE",
		"\U0001F603",
		"Hi! \U0001F601",
		"\U0001F383\U0001F47F\U0001F480\U0001F47B",
	}
)

// String returns an itertor over string, with "good" instances of the string type.
//
// These could be useful when writing tests.
//
// Example
//
//	iterator := goodinstances.String()
//	
//	defer iterator.Close()
//	
//	for iterator.Next() {
//		var datum string
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
func String() Iterator {

	slice := append([]string(nil), sliceString...)
	for i:=0; i<10; i++ {
		s := fmt.Sprintf("%x", randomness.Int63())
		slice = append(slice, s)
	}

	iterator := iterstring.Slice{
		Slice: slice,
	}

	return &iterator
}
