package goodinstances

import (
	"github.com/reiver/go-iter/time"

	"time"
)

const (
	durationDayish  = time.Hour * time.Duration(24)
	durationYearish = time.Hour * time.Duration(24 * 365)
)

var (
	sliceTime = []time.Time{
		time.Date(1976,time.January,1, 10,25,5,625, time.UTC),
		time.Now(),
	}
)

// Time returns an itertor over time.Time, with "good" instances of the time.Time type.
//
// These could be useful when writing tests.
//
// Example
//
//	iterator := goodinstances.Time()
//	
//	defer iterator.Close()
//	
//	for iterator.Next() {
//		var datum time.Time
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
func Time() Iterator {

	slice := append([]time.Time(nil), sliceTime...)
	for i:=0; i<20; i++ {
		datum := slice[len(slice)-1].
			Add(-1 * durationYearish  * time.Duration(randomness.Intn(    50))).
			Add(-1 * durationDayish   * time.Duration(randomness.Intn(   100))).
			Add(-1 * time.Hour        * time.Duration(randomness.Intn(   500))).
			Add(-1 * time.Minute      * time.Duration(randomness.Intn(  1000))).
			Add(-1 * time.Second      * time.Duration(randomness.Intn(  5000))).
			Add(-1 * time.Millisecond * time.Duration(randomness.Intn( 10000))).
			Add(-1 * time.Microsecond * time.Duration(randomness.Intn( 50000))).
			Add(-1 * time.Nanosecond  * time.Duration(randomness.Intn(100000)))

		slice = append(slice, datum)
	}

	iterator := itertime.Slice{
		Slice: slice,
	}

	return &iterator
}
