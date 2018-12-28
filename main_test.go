package main

import (
	"testing"
)

// this function tests term generation.
func TestGenTerm(t *testing.T) {
	// Let's make an slice of example
	// terms for testing
	terms := make([]Term, 0)
	for i := 0; i < 150; i++ {
		terms = append(terms, GenTerm(1))
	}

	// First check: any of the terms generated are blank?
	for _, v := range terms {
		if v.formatTerm() == "" {
			// welp we found what we were looking for
			t.Error("Term was found blank, term -->", v.formatTerm())
			break
		}
	}

	// Second check: any of the terms generated don't make sense?
	// More specifically: are they obeying the rules of formatting a term?
	for _, v := range terms {
		if v.Coefficient == 0 {
			// no null coefficients please
			t.Error("Coefficient was found to be zero, term -->", v.formatTerm())
			break
		}
		if v.Coefficient == 1 {
			// Verify that the coefficient was displayed
			// According to my math textbook, it should not be
			// Note: have to convert to string, because it seems
			// indexing a go string returns a byte
			if string(v.formatTerm()[0]) == "1" {
				// if the coefficient is 1 it should only display the variable
				t.Error("Coefficient was displaying as 1, when it should not have been displaying at all. term -->", v.formatTerm())
			}
		}
		if v.Degree > 1 {
			// When initiating the term generator, we specified that
			// the limit for the degree *must* be 1. If it varies over
			// that it would break the structure of the equation.
			// Linear equations are called linear because for them
			// the max degree as whole for the equation is 1. If it
			// varies it's not a linear equation anymore.
			t.Error("Degree not according to parameters, term -->", v.formatTerm())
		}
	}
}
