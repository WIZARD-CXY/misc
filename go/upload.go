package main

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"qiniupkg.com/api.v7/kodo"
)

// qiniu upload file to bucket
func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE:", os.Args[0], "fileName")
		os.Exit(1)
	}
	AK := "xxxx"
	SK := "xxxx"
	bucketName := "xxxx"
	kodo.SetMac(AK, SK)

	zone := 0
	c := kodo.New(zone, nil)
	bucket := c.Bucket(bucketName)

	ctx := context.Background()

	localFile := os.Args[1]
	err := bucket.PutFile(ctx, nil, localFile, localFile, nil)

	if err != nil {
		fmt.Println("fail to upload")
		os.Exit(1)
	}

	fmt.Println(localFile + " uploaded")
}
