/*
Package goodinstances provides tool for getting good instances of various thing; this is useful for writing tests.

Example

	import (
		"github.com/reiver/go-goodinstances"
	
		"fmt"
		"testing"
	)
	
	// ...
	
	func TestSomething(t *testing.T) {
	
		tests := []struct{
			Int64    int64
			Expected string
		}{}
	
		// Create some "good" tests.
		func() {
			iterator := goodinstances.Int64()
	
			defer iterator.Close()
	
			for iterator.Next() {
				var datum int64
	
				if err := iterator.Decode(&datum); nil != err {
					return err
				}
	
				test := struct{
					Int64    int64
					Expected string
				}{
					Int64:    datum,
					Expected: fmt.Sprintf("Something(%d)", datum),
				}
	
				tests = append(tests, test)
			}
			if err := iterator.Err(); nil != err {
				return err
			}
		}()
	
	
		for testNumber, test := range tests {
	
			//@TODO
	
		}
	}

*/
package goodinstances
