package main

import (
	"reflect"
	"testing"
)

/*
TestPolynomial tests the polynomial generator. Its first test runs the generator, verifies the result,
then converts the output into a list of terms, then converts the list of terms back into a string, then
verifies it. The result should be the same. If not, the method Fail() of "t" is called, which ends the process.
*/
func TestPolynomial(t *testing.T) {
	p := NewPolynomial(4)
	termSlice := make([]Term, 0)
	termSlice = p.ToTerms()
	if len(termSlice) > 4 {
		t.Error("Length of terms slice is more than needed")
		t.Fail()
	}
	bp := PolynomialFromTerms(termSlice)
	if !reflect.DeepEqual(bp, p) {
		t.Error("Polynomials are not equal")
		t.Fail()
	}
}
