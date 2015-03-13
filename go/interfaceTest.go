package main

import "fmt"

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type User struct {
	id   int
	name string
}

func (user *User) Print() {
	fmt.Println(user.String())
}

func (user *User) String() string {
	return fmt.Sprintf("haha id is %d, name is %s", user.id, user.name)
}

func Print(v interface{}) {
	fmt.Printf("%T , %v\n", v, v)
}

func main() {
	u := &User{1, "hsm"}
	u.Print()
	Print(u)
	Print(1)

}
