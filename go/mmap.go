package main

import (
    "fmt"
    "os"
    "syscall"
    "unsafe"
    "encoding/binary"
)

func main() {
    const n = 10
    t := int(unsafe.Sizeof(0)) * n

    map_file, err := os.Create("/tmp/test.dat")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    _, err = map_file.Seek(int64(t-1), 0)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    _, err = map_file.Write([]byte(" "))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    mmap, err := syscall.Mmap(int(map_file.Fd()), 0, int(t), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    map_array := (*[n]int)(unsafe.Pointer(&mmap[0]))

    for i := 0; i < n; i++ {
        map_array[i] = i * i
    }
    fmt.Println("before",*map_array)

     _, err = map_file.Seek(8, 0)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    buf := make([]byte, binary.MaxVarintLen64)
    n1 := binary.PutUvarint(buf,127)
    fmt.Println("nimei",n)
    _, err = map_file.Write(buf[:n1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("after",*map_array)
    map_array[0]=10
    err = syscall.Munmap(mmap)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    mmap, err = syscall.Mmap(int(map_file.Fd()), 0, int(t), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
    map_array = (*[n]int)(unsafe.Pointer(&mmap[0]))
    //fmt.Println(*map_array)

    err = syscall.Munmap(mmap)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    err = map_file.Close()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
