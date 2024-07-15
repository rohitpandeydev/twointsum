//This package takes two  number of type Intger or Float [fromhere] and returns sum of them
// [fromhere]: https://golang.org/x/exp/constraints

package twointsum

import (
	constra "golang.org/x/exp/constraints"
)

type Number interface {
	constra.Float | constra.Integer
}

// Add takes two number as parameter and return an number which is sum of both the input parameter same as [numberadditon]
// [numberaddition]:  https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
