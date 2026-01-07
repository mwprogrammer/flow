/*
Package types provides type-related utility functions for the Flow library.

Author: Chisomo Chiweza (mwprogrammer)
*/
package types

// IsString checks that a value is a string
func IsString(i any) bool {

	_, ok := i.(string)

	return ok

}
