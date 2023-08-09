package exampleCode

import "math/rand"

func ReadNumber() int {
	rnr := 10

	return rand.Intn(rnr)
}
