package goodinstances

import (
	"fmt"
	"time"

	"testing"
)

// This code is kept around so can dump and example of the funcs running.
//
// The way it is, it will NOT run during a call to "gb test".
//
// To make it run (during a call to "gb test") rename it from "fTestSomething" to "TestSomething".
// (I.e., remove the "f" at the beginning of the name.)
func fTestSomething(t *testing.T) {

	func(){
		iterator := Bool()
		defer iterator.Close()

		for iterator.Next() {
			var datum bool

			if err := iterator.Decode(&datum); nil != err {
				panic(err)
			}

			fmt.Printf("(%T) %v\n", datum, datum)
		}
		if err := iterator.Err(); nil != err {
			panic(err)
		}
	}()

	func(){
		iterator := Float64()
		defer iterator.Close()

		for iterator.Next() {
			var datum float64

			if err := iterator.Decode(&datum); nil != err {
				panic(err)
			}

			fmt.Printf("(%T) %v\n", datum, datum)
		}
		if err := iterator.Err(); nil != err {
			panic(err)
		}
	}()

	func(){
		iterator := Int64()
		defer iterator.Close()

		for iterator.Next() {
			var datum int64

			if err := iterator.Decode(&datum); nil != err {
				panic(err)
			}

			fmt.Printf("(%T) %v\n", datum, datum)
		}
		if err := iterator.Err(); nil != err {
			panic(err)
		}
	}()

	func(){
		iterator := String()
		defer iterator.Close()

		for iterator.Next() {
			var datum string

			if err := iterator.Decode(&datum); nil != err {
				panic(err)
			}

			fmt.Printf("(%T) %v\n", datum, datum)
		}
		if err := iterator.Err(); nil != err {
			panic(err)
		}
	}()

	func(){
		iterator := Time()
		defer iterator.Close()

		for iterator.Next() {
			var datum time.Time

			if err := iterator.Decode(&datum); nil != err {
				panic(err)
			}

			fmt.Printf("(%T) %v\n", datum, datum)
		}
		if err := iterator.Err(); nil != err {
			panic(err)
		}
	}()

	t.Errorf("WOW!")
}
