package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
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

func genTerm(degree int) Term {
	rand.Seed(time.Now().UnixNano())
	cof := rand.Intn(10)
	deg := degree
	if degree == 2 {
		if randbool() == false {
			deg = 1
		}
	}
	if cof == 0 && deg == 1 {
		rand.Seed(time.Now().UnixNano())
		cof = rand.Intn(10)
		if cof == 0 && deg == 1 {
			rand.Seed(time.Now().UnixNano())
			cof = rand.Intn(10)
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
	if answer == "" {
		t.Degree += 1
		return t.formatTerm()
	}
	return answer
}

// Polynomial represents
// a random polynomial problem
type Polynomial struct {
	Terms []Term
}

// NewPolynomial generates a completely
// random Polynomial
func NewPolynomial(terms int) Polynomial {
	p := Polynomial{make([]Term, 0)}
	for i := 1; i < terms; i++ {
		rand.Seed(time.Now().UnixNano())
		x := genTerm(rand.Intn(7))
		p.Terms = append(p.Terms, x)
	}
	return p
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

// NewLinearEquation generates a completely random
// LinearEquation
func NewLinearEquation() LinearEquation {
	t := make([]Term, 0)
	operand := "+"
	rand.Seed(time.Now().UnixNano())
	t = append(t, genTerm(rand.Intn(1)))
	if t[0].Degree == 1 {
		rand.Seed(time.Now().UnixNano())
		t = append(t, genTerm(rand.Intn(1)))
	} else {
		t = append(t, genTerm(1))
	}
	t = append(t, genTerm(0))
	return LinearEquation{t, operand}
}

func (l LinearEquation) Solve() {
	ls := l.FormatEquation()
	// 12x - 4 = 20
	// 12x - 4 + 4 = 20 + 4
	// 12x = 20 + 4

	// identify the left and right sides of the equation
	var leftEq, rightEq string
	eqArr := strings.FieldsFunc(ls, func(c rune) bool { return c == '=' })
	leftEq = eqArr[0]
	rightEq = eqArr[1]

	// spread over the operations
	leftEqSlice := strings.Split(leftEq, " ")
	for i, v := range leftEqSlice {
		if v == "+" || v == "-" {
			numStr := leftEqSlice[i+1]
			var strToSuffix string
			if v == "+" {
				strToSuffix = "-" + numStr
			} else if v == "-" {
				strToSuffix = "+" + numStr
			}
			leftEq += strToSuffix
			rightEq += strToSuffix
		}
	}
}

// FormatEquation formats the linear equation
func (l LinearEquation) FormatEquation() string {
	return fmt.Sprintf("%s %s %s = %s", l.Terms[0].formatTerm(), l.operand, l.Terms[1].formatTerm(), l.Terms[2].formatTerm())
}

type QuadracticEquation struct {
	Terms   []Term
	operand string
}

// NewQuadracticEquation generates a completely random
// QuadracticEquation
func NewQuadracticEquation() QuadracticEquation {
	t := make([]Term, 0)
	operand := "+"
	t = append(t, genTerm(2))
	t = append(t, genTerm(2))
	t = append(t, genTerm(0))
	return QuadracticEquation{t, operand}
}

// FormatEquation formats the quadractic equation
func (q QuadracticEquation) FormatEquation() string {
	return fmt.Sprintf("%s %s %s = %s", q.Terms[0].formatTerm(), q.operand, q.Terms[1].formatTerm(), q.Terms[2].formatTerm())
}

// Question represents client request for problems.
type Question struct {
	// Grade    int    `json:"grade"`
	// Syllabus string `json:"syllabus"`
	// Mode     string `json:"mode"`
	Pattern string `json:"pattern"`
	Amount  int    `json:"amount"`
}

// RequestReply represents a server reply made
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
	if q.Pattern == "lineareq" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewLinearEquation().FormatEquation())
		}
		return rp
	}
	if q.Pattern == "quadractic" || q.Pattern == "Quadractic" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewQuadracticEquation().FormatEquation())
		}
		return rp
	}
	return nil
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "{\"msg\": \"No such URL\", \"code\": 404}")
			return
		}
		w.Header().Set("Content-Type", "text/html")
		b, err := ioutil.ReadFile("frontend/index.html")

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "{\"msg\": \"Error while trying to serve index: %v\", \"code\": 500}", err)
			return
		}
		w.Write(b)
	})
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// grade, err := strconv.Atoi(q.Get("grade"))
		// if err != nil {
		// 	fmt.Fprintf(w, "Error while processing query (the grade part): %v", err)
		// 	return
		// }
		// syllabus := q.Get("syllabus")
		// mode := q.Get("mode")
		pattern := q.Get("pattern")
		amount, err := strconv.Atoi(q.Get("amount"))
		if err != nil {
			fmt.Fprintf(w, "{\"msg\": \"Error while processing query (the amount part) because: %v\", \"code\": 500}", err)
			return
		}
		// turn the request into something that can be processed
		// lq := Question{grade, syllabus, mode, pattern, amount}
		lq := Question{pattern, amount}
		if lq.Reply() == nil {
			fmt.Fprintf(w, "{\"msg\": \"URL not supported\", \"code\": 501}")
			return
		}
		rq := &RequestReply{lq, lq.Reply()}
		b, err := json.Marshal(rq)
		if err != nil {
			fmt.Fprintf(w, "{\"msg\": \"Can't marshal the answer because: %v\", \"code\": 500}", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
	http.HandleFunc("/api/pdf", func (w http.ResponseWriter, r *http.Request) {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Times", "", 16)
		pdf.Cell(28, 10, "1. Solve for")
		pdf.SetFont("Times", "I", 16)
		pdf.Cell(3, 10, "x")
		pdf.SetFont("Times", "", 16)
		pdf.Cell(-20, 10, ":")
		// make a for loop with lots of numbers representing the problems
		// pdf.Cell(0.1, 40, "1. 12x * 3 = 4")
		questions := make([]string, 0)
		for i := 1; i < 10; i++ {
			questions = append(questions, NewLinearEquation().FormatEquation())
		}
		alphabet := "abcdefghijklmnopqrstuvwxyz"
		var startPos float64 = 30
		for i, v := range questions {
			pdf.Cell(0.1, startPos+10, fmt.Sprintf("%s. %s __________", string(alphabet[i]), v))
			startPos += 18
		}
		w.Header().Set("Content-Type", "application/pdf")
		err := pdf.Output(w)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "{\"msg\": \"Error while processing PDF document: %v\", \"code\": 500}", err)
		}
	})
	fmt.Println("Starting API server.")
	http.ListenAndServe(":8080", nil)
}
