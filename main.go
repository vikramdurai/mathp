package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// Term represents an individual
// term in any algebraic pattern
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

func genTerm(degree bool) Term {
	rand.Seed(time.Now().UnixNano())
	cof := rand.Intn(10)
	deg := 1
	if degree {
		rand.Seed(time.Now().UnixNano())
		deg = rand.Intn(10)
	}
	if cof == 0 && deg == 1 {
		rand.Seed(time.Now().UnixNano())
		deg = rand.Intn(10)
		if cof == 0 && deg == 1 {
			rand.Seed(time.Now().UnixNano())
			deg = rand.Intn(10)
		}
	}
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

// Polynomial represents
// a random polynomial problem
type Polynomial struct {
	Terms []Term
}

// LinearEquation represents a
// random linear equation to solve
type LinearEquation struct {
	// The linear equation looks
	// like the following:
	// N\Nx o Nx\N = N o Nx\X
	// N  means a constant coefficent.
	// Nx means a term with coefficent and variable
	// X  means a constant variable.
	// o  means either addition or subtraction.
	// A regular expression to validate a Linear Equation
	// [123456789]|([123456789][abcdefghijklmnop]) [-+] = [123456789] [-+] ([123456789][abcdefghijklmnop])|[abcdefghijklmnop]
	Terms   []Term
	operand string
}

// NewPolynomial generates a completely
// random Polynomial
func NewPolynomial(terms int) Polynomial {
	p := Polynomial{make([]Term, 0)}
	for i := 1; i < terms; i++ {
		x := genTerm(true)
		p.Terms = append(p.Terms, x)
	}
	return p
}

// NewLinearEquation generates a completely random
// LinearEquation
func NewLinearEquation() LinearEquation {
	t := make([]Term, 0)
	restr := "[123456789]|([123456789][abcdefghijklmnop]) [-+] [123456789]|([123456789][abcdefghijklmnop]) = [123456789] [-+] ([123456789][abcdefghijklmnop])|[abcdefghijklmnop]"
	re := regexp.MustCompile(restr)
	operand := ""
	if randbool() == true {
		operand = "+"
	} else {
		operand = "-"
	}
	for i := 1; i < 6; i++ {
		x := genTerm(false)
		t = append(t, x)
	}
	// now we need to verify
	lestr := fmt.Sprintf("%s %s %s = %s %s %s", t[0].formatTerm(), operand, t[1].formatTerm(), t[2].formatTerm(), operand, t[3].formatTerm())
	if !re.MatchString(lestr) {
		return NewLinearEquation()
	}
	return LinearEquation{t, operand}
}

// FormatEquation formats the linear equation
func (l LinearEquation) FormatEquation() string {
	t := l.Terms
	return fmt.Sprintf("%s %s %s = %s %s %s", t[0].formatTerm(), l.operand, t[1].formatTerm(), t[2].formatTerm(), l.operand, t[3].formatTerm())
}

// ToTerms converts a Polynomial to
// a slice of Terms
func (p Polynomial) ToTerms() []Term {
	return p.Terms
}

// PolynomialFromTerms creates a Polynomial
// from a slice of Terms
func PolynomialFromTerms(t []Term) Polynomial {
	return Polynomial{t}
}

// Format formats the polynomial
// and pulls it together
func (p Polynomial) Format() string {
	polynomial := ""
	for i, v := range p.Terms {
		polynomial += v.formatTerm()
		if i+1 < len(p.Terms) {
			if p.Terms[i+1].Positive == true {
				polynomial += "+"
				continue
			}
			if p.Terms[i+1].Positive != true {
				continue
			}
		}
	}
	return polynomial
}

// Question represents client request for problems.
type Question struct {
	Grade    int    `json:"grade"`
	Syllabus string `json:"syllabus"`
	Mode     string `json:"mode"`
	Pattern  string `json:"pattern"`
	Amount   int    `json:"amount"`
}

// RequestReply represents a reply made
// to a client request
type RequestReply struct {
	Request Question `json:"request"`
	Reply   []string `json:"reply"`
}

// Reply creates problems according to question criteria
// and returns it.
func (q Question) Reply() []string {
	if q.Pattern == "polynomial" || q.Pattern == "Polynomial" {
		rp := []string{}
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewPolynomial(rand.Intn(3)+1).Format())
		}
		return rp
	}
	if q.Pattern == "linearequation" || q.Pattern == "Linearequation" || q.Pattern == "LinearEquation" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewLinearEquation().FormatEquation())
		}
		return rp
	}
	return nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile("index.html")
		if err != nil {
			w.Write([]byte("Error while trying to serve index"))
		}
		w.Write(b)
	})
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		grade, err := strconv.Atoi(q.Get("grade"))
		if err != nil {
			fmt.Fprintf(w, "Error while processing query (the grade part): %v", err)
			return
		}
		syllabus := q.Get("syllabus")
		mode := q.Get("mode")
		pattern := q.Get("pattern")
		amount, err := strconv.Atoi(q.Get("amount"))
		if err != nil {
			fmt.Fprintf(w, "Error while processing query (the amount part): %v", err)
			return
		}
		// turn the request into something that can be processed
		lq := Question{grade, syllabus, mode, pattern, amount}
		if lq.Reply() == nil {
			fmt.Fprintf(w, "Sorry, we don't support that kind of problem yet")
			return
		}
		rq := &RequestReply{lq, lq.Reply()}
		b, err := json.Marshal(rq)
		if err != nil {
			fmt.Fprintf(w, "Can't marshal the answer. Sorry!")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
	fmt.Println("Starting API server.")
	http.ListenAndServe(":8080", nil)
}
