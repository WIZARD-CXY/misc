package main 

import (
    "fmt"
    "unsafe"
)

type Slice struct{
    ptr unsafe.Pointer
    len int
    cap int
}

var intlen = int(unsafe.Sizeof(int(0)))

func main(){
    s := make([]int,10,20)

    if intlen ==4{
        m:= *(*[4+4*2]byte)(unsafe.Pointer(&s))
        fmt.Println("4",m)
    }else{
        m := *(*[8+8*2]byte)(unsafe.Pointer(&s))
        fmt.Println("8",m)
    }

    slice := (*Slice)(unsafe.Pointer(&s))
    fmt.Println("slice struct", slice)

    fmt.Printf("ptr:%v len:%v cap:%v\n",slice.ptr,slice.len,slice.cap)

    fmt.Println("golang slice len and cap",len(s),cap(s))

    s[0]=0
    s[1]=1
    s[2]=2

    arr := *(*[3]int)(unsafe.Pointer(slice.ptr))
    fmt.Printf("array %p %p %p\n",&arr,&s,slice.ptr)
    fmt.Println("array",arr)
    arr[0]=3
    fmt.Println("array",arr,s)
    fmt.Printf("array %p %p\n",&arr,&s)

    slice.len=15

    fmt.Println("slice len", slice.len)
    fmt.Println("golang slice len",len(s))
}