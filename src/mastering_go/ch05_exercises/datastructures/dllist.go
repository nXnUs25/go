package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

type Node struct {
	value interface{}
	prev  *Node // perv null mean last node
	next  *Node // next null mean first node
}

type DLL struct {
	size  int
	first *Node
	last  *Node
}

type SDLL struct {
	size int
	node *Node
}

func main() {
	debug = true
	list := &DLL{}
	counter := 0
	pdf("Empty?: [%v]\n", list.IsEmpty())
	pdf("first: [%v], size:[%v]\n", list.First(), list.Length())
	pdf("node first data: %v, node last data: %v\n", list.FirstValue(), list.LastValue())
	pdf("node last: %v\n", list.Last())
	for i := 1; i < 4; i++ {
		counter++
		list.InsertrBeginning(i)
		pdf("%v Empty?: %v\n", counter, list.IsEmpty())
		pdf("%v node:%v, size:%v\n", counter, list.First(), list.Length())
		pdf("%v node data: %v\n", counter, list.Last())
		pdf("node first data: %v, node last data: %v\n", list.FirstValue(), list.LastValue())
		pdf("node last: %v\n", list.Last())
	}
}

func (dll *DLL) InsertrBeginning(v interface{}) bool {
	newNode := &Node{v, nil, nil}
	if dll.IsEmpty() {
		dll.first = newNode
		dll.last = newNode
	} else {
		newNode.next = dll.first
		dll.first.prev = newNode
		dll.first = newNode
	}
	dll.size++
	return true
}

func (dll *DLL) IsEmpty() bool {
	return dll.first == nil
}

func (dll *DLL) Length() int {
	return dll.size
}

func (dll *DLL) First() *Node {
	if dll.first != nil {
		return dll.first
	}
	return nil
}

func (dll *DLL) FirstValue() interface{} {
	if dll.first != nil && dll.first.value != nil {
		return dll.first.value
	}
	return nil
}

func (dll *DLL) Last() *Node {
	if dll.last != nil {
		return dll.last
	}
	return nil
}

func (dll *DLL) LastValue() interface{} {
	if dll.last != nil && dll.last.value != nil {
		return dll.last.value
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
