package main

import "fmt"

// some interface
type Stringer interface {
	String() string
}

// a struct that implements Stringer interface
type Struct1 struct {
	field1 string
}

func (s Struct1) String() string {
	return s.field1
}

// another struct that implements Stringer interface, but has a different set of fields
type Struct2 struct {
	field1 []string
	dummy  bool
}

func (s Struct2) String() string {
	return fmt.Sprintf("%v, %v", s.field1, s.dummy)
}

// container that can embedd any struct which implements Stringer interface
type StringerContainer struct {
	Stringer
	haha
	zoom int
}

func test(Ia I) {
	Ia.sayhaha()
	//Ia.(StringerContainer).zoom = 11
	fmt.Println(Ia)

}

// you can change the default sayxixi implementation
func (StringerContainer) sayxixi() {
	fmt.Println("xixihaha")
}

type haha interface {
	sayhaha()
	sayxixi()
}
type S struct {
}

func (S) sayhaha() {
	fmt.Println("sayhaha")
}
func (S) sayxixi() {
	fmt.Println("sayxixi")
}

type I interface {
	haha
	Stringer
}

func NewContainer() I {
	return &StringerContainer{Struct1{"This is Struct1"}, S{}, 1}
}

func main() {
	// the following prints: This is Struct1
	aa := NewContainer()
	aa.sayhaha()
	aa.sayxixi()
	fmt.Println(aa)

	test(aa)

	fmt.Println(aa.(*StringerContainer).zoom)

	// the following prints: [This is Struct1], true
	//fmt.Println(StringerContainer{Struct2{[]string{"This", "is", "Struct1"}, true}})
	// the following does not compile:
	// cannot use "This is a type that does not implement Stringer" (type string)
	// as type Stringer in field value:
	// string does not implement Stringer (missing String method)
	//fmt.Println(StringerContainer{"This is a type that does not implement Stringer"})
}
