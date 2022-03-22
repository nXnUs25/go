package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

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

type Node struct {
	value interface{}
	next  *Node // one before the tail
}

type FIFO struct {
	size int
	node *Node // taiil
}

func (f *FIFO) IsEmpty() bool {
	return f.node == nil
}

func (f *FIFO) Length() int {
	return f.size
}

func (f *FIFO) HasNext() bool {
	return f.node != nil && f.node.next != nil
}

func (f *FIFO) getNode() *Node {
	if f.node != nil {
		return f.node
	}
	return nil
}

func (f *FIFO) getNodeValue() interface{} {
	if f != nil && f.node != nil {
		return f.node.value
	}
	return nil
}

func (f *FIFO) getNodeNext() *Node {
	if f != nil && f.node != nil && f.node.next != nil {
		return f.node.next
	}
	return nil
}

func (n *Node) getNext() *Node {
	if n != nil && n.next != nil {
		return n.next
	}
	return nil
}

func (n *Node) getValue() interface{} {
	if n != nil {
		return n.value
	}
	return nil
}

/*
Make the necessary changes to the code of queue.go in order to store floatingpoint numbers instead of integers.

Personal change will create a new queue which can store all types base on empty interface.
*/
func main() {
	expected := `
			we should see this output:
			where results should be read counter.value
		push 1 2 3 4
		pop 1.1 2.2 3.3
		left 4
		push 0 1 2 3 4
		pop 4.4 5.0 6.1
		left 2 3 4
		push a b c d
		pop 7.2 8.3 9.4 10.a
		left b c d
		push "queue is done."
		pop(all) 11.b 12.c 13.d 14."queue is done."
		size 0 and empty
	`

	pln(expected)
	pln("\nOutput:")

	counter := 1
	fifo := &FIFO{}
	pln(fifo.Traverse())
	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)
	fifo.Push(4)
	pln("\n" + fifo.Traverse())
	v, ok := fifo.Pop()
	if ok {
		pf("%v - Node Pop value: %v -- %[1]v.%[2]v\n", counter, v)
		counter++
	}
	v, ok = fifo.Pop()
	if ok {
		pf("%v - Node Pop value: %v -- %[1]v.%[2]v\n", counter, v)
		counter++
	}
	v, ok = fifo.Pop()
	if ok {
		pf("%v - Node Pop value: %v -- %[1]v.%[2]v\n", counter, v)
		counter++
	}
	pln(fifo.Traverse())
	for i := 0; i < 5; i++ {
		fifo.Push(i)
	}
	pln("\n" + fifo.Traverse())

	for i := 1; i < 4; i++ {
		v, ok := fifo.Pop()
		pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, i, v, ok)
		counter++
	}
	pln(fifo.Traverse())

	for _, c := range "abcd" {
		fifo.Push(string(c))
	}
	pln("\n" + fifo.Traverse())

	for i := 0; i < 4; i++ {
		v, ok := fifo.Pop()
		pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, i, v, ok)
		counter++
	}
	pln(fifo.Traverse())
	fifo.Push("queue is done.")
	pln("\n" + fifo.Traverse())
	for i := 0; i < 4; i++ {
		v, ok := fifo.Pop()
		pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, i, v, ok)
		counter++
	}
	pln(fifo.Traverse())

	pf("Empty?: %v\n", fifo.IsEmpty())
	pf("node:%v, size:%v\n", fifo.getNode(), fifo.Length())
	pf("node data: %v, node next: %v\n", fifo.getNodeValue(), fifo.getNodeNext())

}

func (f *FIFO) Push(v interface{}) bool {
	if f.IsEmpty() {
		f.node = &Node{v, nil}
		f.size++
		return true
	}
	next := &Node{v, f.node}
	f.node = next
	f.size++
	return true
}

func (f *FIFO) Pop() (interface{}, bool) {
	if f.IsEmpty() {
		return nil, false
	}
	if f.Length() == 1 {
		v := f.getNodeValue()
		f.node = nil
		f.size--
		return v, true
	}
	tail := f.node
	temp := f.node
	for f.HasNext() {
		temp = f.node
		f.node = f.node.next
	}
	v := (temp.next).value
	temp.next = nil

	f.node = tail
	f.size--

	return v, true
}

func (f *FIFO) Traverse() (queue string) {
	if f.Length() == 0 {
		queue = "Empty Queue."
		return
	}

	tail := f.node
	for f.node != nil {
		if f.node != tail {
			queue = fmt.Sprintf("%v->", f.node.value) + queue
		} else {
			queue += fmt.Sprintf("%v", f.node.value)
		}
		f.node = f.node.next
	}
	f.node = tail
	queue = "head: " + queue + " :tail"
	return
}
