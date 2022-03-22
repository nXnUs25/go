package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

type Node struct {
	value interface{}
	next  *Node
}

var (
	size      = 0
	queue     = new(Node)
	NodeError = errors.New("Node is not valid.")
	p         = fmt.Print
	pf        = fmt.Printf
	pln       = fmt.Println
	l         *log.Logger
	I         = "[INFO] "
	D         = "[DEBUG] "
	E         = "[ERROR] "
	lf        = log.Printf
	lln       = log.Println
	pwd, _    = os.Getwd()
	debug     = true
	buf       bytes.Buffer
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

func (n *Node) SetValue(v interface{}) error {
	if n == nil {
		return NodeError
	}
	n.value = v
	return nil
}

func (n *Node) GetValue() interface{} {
	if n != nil {
		return n.value
	}
	return nil
}

func CreateNode(v interface{}) *Node {
	n := &Node{}
	n.SetValue(v)
	return n
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(node *Node) error {
	if n == nil {
		return NodeError
	}
	n.next = queue
	queue = n
	size++
	return nil
}

func (n *Node) Push(v interface{}) bool {
	if queue == nil {
		queue = CreateNode(v)
		size++
		return true
	}
	n2 := CreateNode(v)
	n2.SetNext(queue)
	return true
}

func (n *Node) Pop() (interface{}, bool) {
	if GetSize() == 0 {
		return nil, false
	}

	if GetSize() == 1 {
		queue = nil
		size--
		return n.GetValue(), true
	}

	temp := n
	for n.HasNext() {
		temp = n
		n = n.GetNext()
	}

	v := temp.GetValue()
	n.SetNext(nil)

	size--
	return v, true
}

func GetSize() int {
	return size
}

func (n *Node) HasNext() bool {
	return n != nil
}

func (n *Node) QueueIterator() {
	if GetSize() == 0 {
		pln("Empty Queue")
	}
	pf(" %4s | %20s |\n", "Node", "Value")
	count := 1
	for n.HasNext() {
		pf(" %4v | %20v |\n", count, n.GetValue())
		n = n.GetNext()
		count++
		if count > size {
			pln("Count:", count)
			break
		}
	}

}

/*
Make the necessary changes to the code of queue.go in order to store floatingpoint numbers instead of integers.

Personal change will create a new queue which can store all types base on empty interface.

****************************************************************
****************************************************************
****************************************************************
*****                                                      *****
*****               BAD DESIGN DO NOT USE IT               *****
*****               BAD DESIGN DO NOT USE IT               *****
*****               BAD DESIGN DO NOT USE IT               *****
*****                                                      *****
****************************************************************
****************************************************************
****************************************************************

*/
func main() {
	l = log.New(&buf, "Queue: ", log.LstdFlags|log.Lshortfile)
	l.SetOutput(&buf)
	queue = nil
	queue.Push(2.2)
	//pln(size, queue.GetValue(), queue.GetNext())
	queue.Push(100)
	//pln(size, queue.GetValue(), queue.GetNext())
	queue.Push("abc")
	pln(size, queue.GetValue(), queue.GetNext())
	queue.Push("abc2")

	queue.QueueIterator()

	v, b := queue.Pop()
	if b {
		pf("Value: %v\n", v)
	}

	queue.QueueIterator()

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

func SetOutput(buf bytes.Buffer) {
	l = log.New(&buf, "Queue: ", log.LstdFlags|log.Lshortfile)
	l.SetOutput(&buf)
}
