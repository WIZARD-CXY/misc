// Created by cgo - DO NOT EDIT

//line /home/wizard/gows/rand.go:1
package rand
//line /home/wizard/gows/rand.go:7

//line /home/wizard/gows/rand.go:6
func Random() int {
	return int(_Cfunc_random())
}
func Seed(i int) {
	_Cfunc_srandom(_Ctype_uint(i))
}
