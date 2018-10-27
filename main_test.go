// package main

// import (
// 	"testing"
// )

// /*
// TestPolynomialGenerator tests the polynomial generator. Its first test runs the generator, verifies the result,
// then converts the output into a list of terms, then converts the list of terms back into a string, then
// verifies it. The result should be the same. If not, the method Fail() of "t" is called, which ends the process.
// */
// func TestPolynomialGenerator(t *testing.T) {
// 	p := PolynomialGenerator(4)
// 	termSlice := make([]Term, 0)
// 	termSlice, err = PolynomialToTerms(p)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	if len(termSlice) > 4 {
// 		t.Fail()
// 	}
// 	bp := TermsToPolynomial(termSlice)
// 	if bp != p {
// 		t.Fail()
// 	}
// }
package main
