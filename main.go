package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// Helper functions.
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

// I just grabbed this code from
// golangcookbook.com/chapters/arrays/reverse
func reverse(numbers []string) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func loopT(n int) []Term {
	termSlice := make([]Term, 0)
	for i := 0; i < n; i++ {
		termSlice = append(termSlice, GenTerm(1))
	}
	return termSlice
}

// Term represents an individual
// term in any algebraic pattern
type Term struct {
	Coefficient int
	Positive    bool
	Variable    string
	Degree      int
}

// GenTerm generates a completely
// random term. This is the backbone
// of the whole program.
func GenTerm(degree int) Term {
	rand.Seed(time.Now().UnixNano())
	cof := rand.Intn(8) + 2
	deg := degree
	if degree == 1 {
		if randbool() == false {
			if randbool() == false {
				deg = 0
			} else {
				deg = 1
			}
		}
	}
	// if cof == 0 {
	// 	GenTerm(deg)
	// }
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
	superscript := []string{"²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹", "¹⁰"}
	answer := ""

	if t.Coefficient != 0 {
		answer += fmt.Sprintf("%d", t.Coefficient)
	}
	if t.Coefficient != 1 {
		if answer != fmt.Sprintf("%d", t.Coefficient) {
			answer += fmt.Sprintf("%d", t.Coefficient)
		}
	}
	if t.Degree >= 1 {
		answer += t.Variable
	}
	if t.Degree > 1 {
		answer += superscript[t.Degree]
	}
	if answer == "" {
		t = GenTerm(t.Degree)
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
func NewPolynomial() Polynomial {
	p := Polynomial{make([]Term, 0)}
	rand.Seed(time.Now().UnixNano())
	numTerms := rand.Intn(3) + 2
	for i := 0; i < numTerms; i++ {
		rand.Seed(time.Now().UnixNano())
		p.Terms = append(p.Terms, GenTerm(rand.Intn(4)+1))
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

// AlgebraExpr represents
// a completely random algebraic
// expression
type AlgebraExpr struct {
	Terms   []Term
	operand string
	// this represents the value of
	// x in any given expression
	VarVal int
}

// random easter egg
var egg = []int{1, 1, 2, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 6, 7, 8, 9, 10}

// NewAlgebraExpr generates a new
// algebraic expression in the form
// of the exported struct AlgebraExpr.
func NewAlgebraExpr() AlgebraExpr {
	rand.Seed(time.Now().UnixNano())
	// this little trick ensures a nonzero
	// value for the variable
	variableValue := rand.Intn(9) + 1
	termSlice := loopT(2)
	var operand string
	if randbool() == false {
		operand = "-"
	} else {
		operand = "+"
	}
	return AlgebraExpr{termSlice, operand, variableValue}
}

// FormatExpr uses the structural representation
// of AlgebraExpr and formats the information to
// a mathematical format
func (alx AlgebraExpr) FormatExpr() string {
	return fmt.Sprintf("%s %s %s (x = %d)", alx.Terms[0].formatTerm(),
		alx.operand, alx.Terms[1].formatTerm(), alx.VarVal)
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
	var operand string
	if randbool() == false {
		operand = "-"
	} else {
		operand = "+"
	}
	t = append(t, GenTerm(1))
	t = append(t, GenTerm(1))
	t = append(t, GenTerm(0))
	return LinearEquation{t, operand}
}

// the unfinished solver code
// func (l LinearEquation) Solve() {
// 	ls := l.FormatEquation()
// 	// 12x - 4 = 20
// 	// 12x - 4 + 4 = 20 + 4
// 	// 12x = 20 + 4

// 	// identify the left and right sides of the equation
// 	var leftEq, rightEq string
// 	eqArr := strings.FieldsFunc(ls, func(c rune) bool { return c == '=' })
// 	leftEq = eqArr[0]
// 	rightEq = eqArr[1]

// 	// spread over the operations
// 	leftEqSlice := strings.Split(leftEq, " ")
// 	for i, v := range leftEqSlice {
// 		if v == "+" || v == "-" {
// 			numStr := leftEqSlice[i+1]
// 			var strToSuffix string
// 			if v == "+" {
// 				strToSuffix = "-" + numStr
// 			} else if v == "-" {
// 				strToSuffix = "+" + numStr
// 			}
// 			leftEq += strToSuffix
// 			rightEq += strToSuffix
// 		}
// 	}
// }

// FormatEquation formats the linear equation
func (l LinearEquation) FormatEquation() string {
	return fmt.Sprintf("%s %s %s = %s", l.Terms[0].formatTerm(), l.operand, l.Terms[1].formatTerm(), l.Terms[2].formatTerm())
}

// QuadracticEquation is mainly
// a second-degree version of
// LinearEquation
type QuadracticEquation struct {
	Terms   []Term
	operand string
}

// NewQuadracticEquation generates a completely random
// QuadracticEquation
func NewQuadracticEquation() QuadracticEquation {
	t := make([]Term, 0)
	var operand string
	if randbool() == false {
		operand = "-"
	} else {
		operand = "+"
	}
	t = append(t, GenTerm(2))
	t = append(t, GenTerm(2))
	t = append(t, GenTerm(0))
	return QuadracticEquation{t, operand}
}

// FormatEquation formats the quadractic equation
func (q QuadracticEquation) FormatEquation() string {
	return fmt.Sprintf("%s %s %s = %s", q.Terms[0].formatTerm(), q.operand, q.Terms[1].formatTerm(), q.Terms[2].formatTerm())
}

// WordProblem represents any random word problem.
// It's a unique class of problem in that it doesn't use
// the GenTerm() function at all, since algebraic terms are
// not strictly related to word problems.
// WordProblem for the time being is merely a disguised arithmetic
// word problem. I'll add more complex ones later. TODO
type WordProblem struct {
	// People represents the characters used in the word problem
	// The number of people is always 2 for the time being. TODO
	People []person
	// e.g apples, oranges
	Item string
	// arithmetic is a must in word problems
	Operator string
}

type person struct {
	name string
	val  int
}

// NewWordProblem makes new Word Problems
func NewWordProblem() WordProblem {
	possibleNames := []string{"Johnny", "Rohit", "Maple", "George", "Ankit", "Eugene", "Vikram"}
	// Remember, we have 2 people with 2 corresponding values
	// these are the names
	rand.Seed(time.Now().UnixNano())
	name1 := possibleNames[rand.Intn(len(possibleNames))]
	rand.Seed(time.Now().UnixNano())
	name2 := possibleNames[rand.Intn(len(possibleNames))]
	// these are the values
	// e.g the number of apples Johnny has
	rand.Seed(time.Now().Unix())
	value1 := rand.Intn(9) + 1
	// e.g Vikram gives 5 apples
	rand.Seed(time.Now().Unix())
	value2 := rand.Intn(value1-1) + 1
	// this is the arithmetic that unites them both
	op := ""
	if randbool() {
		op = "+"
	} else {
		op = "-"
	}
	// item of random choice
	rand.Seed(time.Now().Unix())
	randItems := []string{"apples", "oranges", "mangoes", "walnuts", "peanuts"}
	// put it all together
	return WordProblem{[]person{person{name1, value1}, person{name2, value2}}, randItems[rand.Intn(len(randItems))], op}
}

// FormatProblem is the actual function that converts
// numbers into an actual word problem
// e.g (Johnny = 4) - (Rohit = 2) ->
// "Johnny has 4 apples, Rohit took 2 apples. How many does Johnny have?"
func (w WordProblem) FormatProblem() string {
	// this is the return value
	if w.Operator == "+" {
		return fmt.Sprintf("%s has %d %s. %s gave %s %d %s. How many %s does %s now have?",
			w.People[0].name, w.People[0].val, w.Item, w.People[1].name, w.People[0].name, w.People[1].val, w.Item, w.Item, w.People[0].name)
	}
	return fmt.Sprintf("%s has %d %s. %s took %d %s. How many %s does %s have?",
		w.People[0].name, w.People[0].val, w.Item, w.People[1].name, w.People[1].val, w.Item, w.Item, w.People[0].name)
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
	if q.Pattern == "polynomial" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewPolynomial().Format())
		}
		return rp
	}
	if q.Pattern == "alx" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewAlgebraExpr().FormatExpr())
		}
		return rp
	}
	if q.Pattern == "lineq" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewLinearEquation().FormatEquation())
		}
		return rp
	}
	if q.Pattern == "quadr" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewQuadracticEquation().FormatEquation())
		}
		return rp
	}
	if q.Pattern == "wrdp" {
		rp := []string{}
		for i := 0; i < q.Amount; i++ {
			rp = append(rp, NewWordProblem().FormatProblem())
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
	http.HandleFunc("/api/pdf", func(w http.ResponseWriter, r *http.Request) {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Times", "", 16)
		pdf.CellFormat(190, 20, "Solve the expressions using the values given.", "", 1, "TC", false, 0, "")
		dataSlice := make([]AlgebraExpr, 0)
		data := make([]string, 0)
		for i := 0; i < 28; i++ {
			dataSlice = append(dataSlice, NewAlgebraExpr())
		}
		for i, v := range dataSlice {
			data = append(data, fmt.Sprintf("%d.  %s", i+1, v.FormatExpr()))
		}
		reverse(data)
		pdf.SetFont("Times", "", 12)
		// WARNING UNREADABLE CODE AHEAD!
		// READ AT YOUR OWN RISK!
		for i := 1; i < 4; i++ {
			for j := 1; j < 4; j++ {
				var x string
				x, data = data[len(data)-1], data[:len(data)-1]
				pdf.CellFormat(50, 0, x, "", 0, "LT", false, 0, "")
			}
			var x string
			x, data = data[len(data)-1], data[:len(data)-1]
			pdf.CellFormat(50, 35, x, "", 1, "LT", false, 0, "")
			for k := 1; k < 4; k++ {
				x, data = data[len(data)-1], data[:len(data)-1]
				pdf.CellFormat(50, 35, x, "", 0, "LT", false, 0, "")
			}
			x, data = data[len(data)-1], data[:len(data)-1]
			pdf.CellFormat(50, 35, x, "", 1, "LT", false, 0, "")
		}
		var x string
		for j := 1; j < 5; j++ {
			x, data = data[len(data)-1], data[:len(data)-1]
			pdf.CellFormat(50, 0, x, "", 0, "LT", false, 0, "")
		}
		err := pdf.Output(w)
		if err != nil {
			fmt.Fprintf(w, "{\"msg\":\"Can't give pdf output because: %v\", \"code\": 500}", err)
			return
		}
	})
	fmt.Println("Starting API server.")
	http.ListenAndServe(":8080", nil)
}
