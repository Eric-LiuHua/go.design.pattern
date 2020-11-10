package main

import (
	"fmt"
	"reflect"
)

type Op interface {
	GetName() string
}
type A struct{}
type B struct{}
type Factory struct{}

func (a *A) GetName() string {
	return reflect.TypeOf(a).String()
}

func (b *B) GetName() string {
	return reflect.TypeOf(b).String()
}

func (f *Factory) CreateOp(name string) Op {
	switch name {
	case "a":
		return new(A)
	case "b":
		return new(B)
	default:
		panic("name no exist!")
	}
}

func main() {
	var f = new(Factory)
	p := f.CreateOp("a")
	fmt.Println(p.GetName())
	p = f.CreateOp("b")
	fmt.Println(p.GetName())
}
