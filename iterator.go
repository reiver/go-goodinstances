package goodinstances

// Iterator represents an iterator.
//
// It is intentionally made to look similar to the iterator that
// *sql.Rows exposes in the "database/sql" package.
//
// With the one difference being that data comes from the iterator
// with the Decode method.
//
// Example
//
//	defer iterator.Close()
//	
//	for iterator.Next() {
//		var datum ??? // <---- The "???" would be replaced with whatever type you actually have.
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
type Iterator interface {
	Close() error
	Decode(interface{}) error
	Err() error
	Next() bool
}
