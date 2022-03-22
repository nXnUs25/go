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

func main() {
	debug = false
	counter := 1
	fifo := &FIFO{}
	pdf("Empty?: %v\n", fifo.IsEmpty())
	pdf("node:%v, size:%v\n", fifo.getNode(), fifo.Length())
	pdf("node data: %v, node next: %v\n", fifo.getNodeValue(), fifo.getNodeNext())
	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)
	fifo.Push(4)
	pln(fifo.Traverse())
	v, ok := fifo.Pop()
	pdf("%v : ok %v\n", v, ok)
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
	pdf("Empty?: %v\n", fifo.IsEmpty())
	pdf("node:%v, size:%v\n", fifo.getNode(), fifo.Length())
	pdf("node data: %v, node next: %v\n", fifo.getNodeValue(), fifo.getNodeNext())
	for i := 0; i < 5; i++ {
		fifo.Push(i)
	}
	pln(fifo.Traverse())
	for i := 1; i < 4; i++ {
		v, ok := fifo.Pop()
		pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, i, v, ok)
		counter++
	}
	pln(fifo.Traverse())
	for _, c := range "abcd" {
		fifo.Push(string(c))
	}
	pln(fifo.Traverse())
	for i := 0; i < 4; i++ {
		v, ok := fifo.Pop()
		pf("%v : %v - Node Pop Value: %v -> %v -- %[1]v.%[3]v\n", counter, i, v, ok)
		counter++
	}
	pln(fifo.Traverse())
	fifo.Push("queue is done.")
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
	n2 := &Node{v, f.node}
	f.node = n2
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
	/*
		temp := f.getNode()
		for f.HasNext() {
			temp = f.getNode()
			f.node = f.getNodeNext()
		}
	*/
	//debug = false
	count := 0
	pdf("0: f.node %v\n", f.node)
	pdf("0: f.node.next [%v]\n", f.node.next)
	tail := f.node
	temp := f.node
	pdf("tmp = %v\n", temp)
	pdln("\nenterin loop")
	for f.HasNext() {
		pdln("%v", count+1)
		pdf("1: f.node %v\n", f.node)
		pdf("1: f.node.next [%v]\n", f.node.next)
		temp = f.node
		pdf("tmp = %v\n", temp)
		pdf("2: f.node %v\n", f.node)
		f.node = f.node.next
		pdf("2: f.node %v\n", f.node)
		count++
		if f.Length() < count {
			pdf("count to high over quque size: %v, %v\n", f.Length(), count)
			break
		}
		pdln("end of %v", count)
	}
	pdln("ending a loop\n")
	pdf("228 f.node: %v\n", f.node)
	pdf("228 tmp = %v\n", temp)
	v := (temp.next).value
	//	v := temp.getNext().getValue()
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
	queue = "head: " + queue + " :tail\n"
	return
}

/*
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
*/
