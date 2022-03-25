package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

type Node struct {
	value interface{}
	next  *Node
}

type LIFO struct {
	size int
	node *Node
}

func main() {
	counter := 0
	debug = false
	lifo := &LIFO{}
	pdf("Empty?: %v\n", lifo.IsEmpty())
	pdf("node:%v, size:%v\n", lifo.getNode(), lifo.Length())
	pdf("node data: %v, node next: %v\n", lifo.getValue(), lifo.getNext())
	pln(lifo.Traverse())
	for i := 1; i < 4; i++ {
		counter++
		lifo.Push(i)
		pdf("%v Empty?: %v\n", counter, lifo.IsEmpty())
		pdf("%v node:%v, size:%v\n", counter, lifo.getNode(), lifo.Length())
		pdf("%v node data: %v, node next: %v\n", counter, lifo.getValue(), lifo.getNext())
	}
	//counter = 0
	pln(lifo.Traverse())
	for i := 1; i < 4; i++ {
		counter++
		v, k := lifo.Pop()
		pdf("%v Empty?: %v\n", counter, lifo.IsEmpty())
		pdf("%v node:%v, size:%v\n", counter, lifo.getNode(), lifo.Length())
		pdf("%v node data: %v, node next: %v\n", counter, lifo.getValue(), lifo.getNext())
		pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, i, v, k)
	}
	pln(lifo.Traverse())
	counter++
	lifo.Push(10)
	lifo.Push(20)
	pln(lifo.Traverse())
	v, k := lifo.Pop()
	pdf("%v Empty?: %v\n", counter, lifo.IsEmpty())
	pdf("%v node:%v, size:%v\n", counter, lifo.getNode(), lifo.Length())
	pdf("%v node data: %v, node next: %v\n", counter, lifo.getValue(), lifo.getNext())
	pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, "00", v, k)
	pln(lifo.Traverse())
	counter++
	lifo.Push(30)
	v, k = lifo.Pop()
	pdf("%v Empty?: %v\n", counter, lifo.IsEmpty())
	pdf("%v node:%v, size:%v\n", counter, lifo.getNode(), lifo.Length())
	pdf("%v node data: %v, node next: %v\n", counter, lifo.getValue(), lifo.getNext())
	pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, "01", v, k)

	pln(lifo.Traverse())
	pdf("Empty?: %v\n", lifo.IsEmpty())
	pdf("node:%v, size:%v\n", lifo.getNode(), lifo.Length())
	pdf("node data: %v, node next: %v\n", lifo.getValue(), lifo.getNext())
}

func (l *LIFO) Push(v interface{}) bool {
	if l.IsEmpty() {
		l.node = &Node{v, nil}
		l.size = 1
		return true
	}
	tmp := &Node{v, l.node}
	l.node = tmp
	l.size++
	return true

}

func (l *LIFO) Pop() (interface{}, bool) {
	if l.Length() == 0 {
		return nil, false
	}

	if l.Length() == 1 {
		l.size = 0
		v := l.getValue()
		l.node = nil
		return v, true
	}

	v := l.getValue()
	l.node = l.getNext()
	l.size--

	return v, true
}

func (l *LIFO) Traverse() (stack string) {
	if l.Length() == 0 {
		pln("Empty Stack (aka: LIFO)")
		return
	}

	tmp := l.getNode()
	for l.HasNode() {
		if l.getNode() != tmp {
			stack = fmt.Sprintf("%v->", l.getValue()) + stack
		} else {
			stack += fmt.Sprintf("%v", l.getValue())
		}
		//pf("%v -> ", l.getValue())
		l.node = l.getNext()
	}
	pln()
	l.node = tmp
	stack = "head: " + stack + " :tail"
	return
}

func (l *LIFO) IsEmpty() bool {
	return l.node == nil
}

func (l *LIFO) Length() int {
	return l.size
}

func (l *LIFO) HasNode() bool {
	return l.node != nil
}

func (l *LIFO) HasNext() bool {
	return l.HasNode() && l.node.next != nil
}

func (l *LIFO) getNode() *Node {
	return l.node
}

func (l *LIFO) getValue() (value interface{}) {
	if l.node != nil {
		return l.node.value
	}
	return nil
}

func (l LIFO) getNext() (node *Node) {
	if l.node != nil && l.node.next != nil {
		return l.node.next
	}
	return nil
}

//

////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	NodeError = errors.New("Node is not valid.")
	p         = fmt.Print
	pf        = fmt.Printf
	pln       = fmt.Println
	l         *log.Logger
	I         = "[INFO] "
	D         = "[DEBUG] "
	E         = "[ERROR] "
	debug     = false
	GetFunc   = func() string {
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
)

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

func pdln(format string, args ...interface{}) {
	if debug {
		pf(D+format+"\n", args...)
	}
}

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////
