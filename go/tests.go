package main

import (
	"fmt"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"time"
)

var rnd = uint32(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999999))

func NewReqId() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], rnd)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

func main() {
	fmt.Println("Hello, playground",NewReqId())
}
