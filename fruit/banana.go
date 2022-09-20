package fruit

import "fmt"

// Banana is a fruit
type Banana struct{}

// String implements fmt.Stringer
func (b Banana) String() string {
	return "banana"
}

// Interface guard, this makes sure Banana implements fmt.Stringer at compile time
var _ fmt.Stringer = (*Banana)(nil)
