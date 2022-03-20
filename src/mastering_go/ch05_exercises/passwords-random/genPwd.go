package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var (
	p       = fmt.Print
	pf      = fmt.Printf
	pln     = fmt.Println
	buf     bytes.Buffer
	l       *log.Logger
	I       = "[INFO] "
	D       = "[DEBUG] "
	E       = "[ERROR] "
	GetFunc = func() string {
		pc, _, _, _ := runtime.Caller(1)
		return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	}
	GetFuncLine = func() string {
		_, _, line, _ := runtime.Caller(1)
		return fmt.Sprintf("%v", line)
	}
	GetFuncDetails = func() string {
		pc, _, line, _ := runtime.Caller(1)
		return fmt.Sprintf("%v:[%+v]", runtime.FuncForPC(pc).Name(), line)
	}
	lf     = log.Printf
	lln    = log.Println
	pwd, _ = os.Getwd()
	debug  = true
)

/*
	embeded type for rand.Rand
	which we can use Intn function on thad field

*/
type myRand struct {
	*rand.Rand
	seed int64
}

func MyRand(seed int64) *myRand {
	pdfln(GetFuncDetails(), "seed passed: [%v]", seed)
	return &myRand{
		Rand: rand.New(rand.NewSource(seed)),
		seed: seed,
	}
}

func UnixNanoRand() *myRand {
	seed := time.Now().UnixNano()
	pdfln(GetFuncDetails(), "seed passed: [%v]", seed)
	return &myRand{
		Rand: rand.New(rand.NewSource(seed)),
		seed: seed,
	}
}

func UnixRand() *myRand {
	seed := time.Now().Unix()
	pdfln(GetFuncDetails(), "seed passed: [%v]", seed)
	return &myRand{
		Rand: rand.New(rand.NewSource(seed)),
		seed: seed,
	}
}

func (r *myRand) String() (rand, seed string) {
	pdfin(GetFunc(), GetFuncLine())
	rand = fmt.Sprintf("Rand: [%v]", r.Rand)
	seed = fmt.Sprintf("Seed: [%v]", r.seed)
	pdfout(GetFunc(), GetFuncLine())
	return
}

func (r *myRand) Range(start, end int) int {
	pdfin(GetFunc(), GetFuncLine())
	random := r.Intn(end-start) + start
	pdfln(GetFuncDetails(), "Random number is [%v] base on formula: [(%[4]v - %[3]v) + %[3]v]", random, start, end)
	pdfout(GetFunc(), GetFuncLine())
	return random
}

/*
a bit different approch to generate passwords
using custom rand type, and at ths same time we will generate passwords
containing only letters and numbers and only capital letters

Make the necessary changes to the code of queue.go in order to store floatingpoint numbers instead of integers.
only upper case letters
	for c := 65; c <= 90; c++ {
		fmt.Printf("Letter: [%v]:[%v]\n", string(c), c)
	}
*/
func main() {
	pdfin(GetFunc(), GetFuncLine())
	lln("Starting passwords generator")
	r := UnixNanoRand()
	MIN := 65
	MAX := 90
	LEN := 8
	p("Password: > ")
	for c := 1; c <= LEN; c++ {
		random := r.Range(MIN, MAX)
		nc := string(byte(random))
		pf("%v", nc)
	}
	pln(" <")
	pln("Done.")
	pln(r.String())
	pdfout(GetFunc(), GetFuncLine())
}

/*
shuold not be used to often and we should avoid using it...
just added this function to test it and examine vars initialization
*/
func init() {
	pdfin(GetFunc(), GetFuncLine())
	l = log.New(&buf, "genPwds: ", log.LstdFlags|log.Lshortfile)
	l.SetOutput(&buf)
	pdfout(GetFunc(), GetFuncLine())
}

func pdfin(function, line string) {
	if debug {
		f := D + "[IN]: func: < %v >:[%v]\n"
		pf(f, function, line)
	}
}

func pdfmsg(funcdetails, format string, args ...interface{}) {
	if debug {
		f := D + "[MSG]: func: < %v > "
		var a []interface{}
		a = append(a, funcdetails)
		a = append(a, args...)
		pf(f+format, a...)
	}
}

func pdfln(funcdetails, format string, args ...interface{}) {
	if debug {
		f := D + "[MSG]: func: < %v > "
		var a []interface{}
		a = append(a, funcdetails)
		a = append(a, args...)
		pf(f+format+"\n", a...)
	}
}

func pdfout(function, line string) {
	if debug {
		f := D + "[OUT]: func: < %v >:[%v]\n"
		pf(f, function, line)
	}
}

func pdf(format string, args ...interface{}) {
	if debug {
		pf(D+format, args...)
	}
}

func ldf(format string, args ...interface{}) {
	if debug {
		lf(D+format, args...)
	}
}
