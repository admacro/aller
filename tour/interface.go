package main

import (
	"fmt"
	"math"
)

// Method set is very important when dealing with interfaces.
// To understand interface, you need to understand method set first.

// An interface type specifies a method set called its interface.
// The interface of an interface type is the method set it specifies.
// An interface is a set of method signatures.

// A variable of interface type can store a value of any type
// with a method set that is any superset of the interface.
// Such a type is said to implement the interface.
// That is, a type implements any interface comprising any
// subset of its methods.
// A type may implement several distinct interfaces, including
// the empty interface `interface{}`.
// Several distinct types may implement the same interface and
// each type may implement some other distinct interfaces.

// A value of interface can hold any value that implements those methods.
// All methods of an interface a must be implemented in v when assigning v to a.

// When assigning v to a, v must have the same type as the receivers in all
// methods of a; in other words, v is the receiver of the methods in a.
// We can also say the type of v, T implements method m when T is same as the
// type of the receiver of method m.

// So, in general, to make things simple, all methods on a given type should
// have either value or pointer receivers, but not a mixture of both.

// Interfaces are implemented implicitly.
// This means a type implements an interface by implementing its methods.
// There is no such explicit declaration as "A implements B".

type Vertexer interface {
	Abs() float64

	// Error on line#26
	// cannot use &v (type *Vertex) as type Vertexer in assignment:
	// *Vertex does not implement Vertexer (missing Scale method)
	// Scale(f float64)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var vter Vertexer
	v := Vertex{3, 4}
	vter = &v
	fmt.Println(vter.Abs())
}
