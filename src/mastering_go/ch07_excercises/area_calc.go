package main

import (
	"fmt"
	"math"
)

type square struct {
	_type shape
	a     float64
}

func (s *square) String() string { return fmt.Sprintf("%s side %v", s._type.String(), s.a) }
func (s *square) area() float64  { return s.a * s.a }
func (s *square) perimeter() (area float64) {
	for i := 0; i < s._type.sides; i++ {
		area += s.a
	}
	return
}
func (s *square) getType() string          { return s._type.name }
func (s *square) getSides() int            { return s._type.sides }
func (s *square) getSideValues() []float64 { return []float64{s.a} }

type rectangle struct {
	_type shape
	a     float64
	b     float64
}

func (r *rectangle) String() string {
	return fmt.Sprintf("%s side %v abd %v", r._type.String(), r.a, r.b)
}
func (r *rectangle) area() float64 { return r.a * r.b }
func (r *rectangle) perimeter() (area float64) {
	return (2 * r.a) + (2 * r.b)
}
func (r *rectangle) getType() string          { return r._type.name }
func (r *rectangle) getSides() int            { return r._type.sides }
func (r *rectangle) getSideValues() []float64 { return []float64{r.a, r.b} }

type shape struct {
	name  string
	sides int
}

func (s *shape) String() string { return fmt.Sprintf("%v Shape", s.name) }

type circle struct {
	_type shape
	r     float64
}

func (c *circle) String() string {
	return fmt.Sprintf("%s radius %v", c._type.String(), c.r)
}
func (c *circle) area() float64 { return math.Pi * (c.r * c.r) }
func (c *circle) perimeter() (area float64) {
	return (2 * math.Pi * c.r)
}
func (c *circle) getType() string          { return c._type.name }
func (c *circle) getSides() int            { return c._type.sides }
func (c *circle) getSideValues() []float64 { return []float64{c.r} }

type sider interface {
	getSideValues() []float64
}

type shaper interface {
	getType() string
	getSides() int
	sider
}

type geometry interface {
	area() float64
	perimeter() float64
	shaper
}

func main() {
	var g geometry = &square{
		_type: shape{
			name:  "square",
			sides: 4,
		},
		a: 5,
	}
	fmt.Println(g)
	fmt.Printf("%v area of sides %d with length of %v is %v\n", g.getType(), g.getSides(), g.getSideValues(), g.area())
	fmt.Printf("%v perimeter of sides %d with length of %v is %v\n", g.getType(), g.getSides(), g.getSideValues(), g.perimeter())

	g = &rectangle{
		_type: shape{
			name:  "rectangle",
			sides: 4,
		},
		a: 5,
		b: 4,
	}
	fmt.Println(g)
	fmt.Printf("%v area of sides %d with length of %v is %v\n", g.getType(), g.getSides(), g.getSideValues(), g.area())
	fmt.Printf("%v perimeter of sides %d with length of %v is %v\n", g.getType(), g.getSides(), g.getSideValues(), g.perimeter())

	g = &circle{
		_type: shape{
			name:  "circle",
			sides: 1,
		},
		r: 5,
	}
	fmt.Println(g)
	fmt.Printf("%v area of radius %d with length of %v is %v\n", g.getType(), g.getSides(), g.getSideValues(), g.area())
	fmt.Printf("%v perimeter of radius %d with length of %v is %v\n", g.getType(), g.getSides(), g.getSideValues(), g.perimeter())
}
