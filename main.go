package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Term represents an individual
// term in a polynomial
type Term struct {
	Coefficient int
	Positive    bool
	Variable    string
	Degree      int
}

func randbool() bool {
	c := make(chan struct{})
	close(c)
	select {
	case <-c:
		return true
	case <-c:
		return false
	}
}

func genTerm() Term {
	rand.Seed(time.Now().UnixNano())
	deg := rand.Intn(10)
	rand.Seed(time.Now().UnixNano())
	cof := rand.Intn(10)
	vars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	rand.Seed(153)
	varn := rand.Intn(25)
	randBool := randbool()
	err := errors.New("nothing happened")
	if !randBool {
		cof, err = strconv.Atoi(fmt.Sprintf("-%d", cof))
		if err != nil {
			println(err)
			panic(nil)
		}
	}
	return Term{cof, randBool, vars[varn], deg}
}

func (t Term) formatTerm() string {
	superscriptToTwenty := []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹",
		"¹⁰"}
	answer := ""

	if t.Coefficient != 0 {
		answer += fmt.Sprintf("%d", t.Coefficient)
	}
	if t.Degree >= 1 {
		answer += t.Variable
	}
	if t.Degree > 1 {
		answer += superscriptToTwenty[t.Degree]
	}
	return answer
}

// PolynomialGenerator generates a random
// polynomial problem
func PolynomialGenerator(terms int) string {
	polynomial := ""
	for i := 1; i <= terms; i++ {
		x := genTerm()
		polynomial += x.formatTerm()
		if i+1 < terms {
			if x.Positive == true {
				polynomial += "+"
				continue
			}
			plusorminus := randbool()
			if plusorminus == true {
				polynomial += "+"
			}
			if plusorminus == false {
				polynomial += "-"
			}
		}
	}
	return polynomial
}
func main() {
	fmt.Println(PolynomialGenerator(3))
}
