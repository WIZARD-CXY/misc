// Created by cgo - DO NOT EDIT

package rand

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int32) { *dst = syscall.Errno(x) }
type _Ctype_long int64

type _Ctype_uint uint32

type _Ctype_void [0]byte

func _Cfunc_random() _Ctype_long
func _Cfunc_srandom(_Ctype_uint) _Ctype_void
