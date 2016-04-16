package main

// qiniu delete one file from bucket
import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"qiniupkg.com/api.v7/kodo"
)

func main() {
	fileName := os.Args[1]
	AK := "xxxx"
	SK := "xxxx"
	bucketName := "xxxx"
	kodo.SetMac(AK, SK)

	zone := 0
	c := kodo.New(zone, nil)
	bucket := c.Bucket(bucketName)

	ctx := context.Background()

	if err := bucket.Delete(ctx, fileName); err != nil {
		fmt.Println("delete failed")
		os.Exit(1)
	}
	fmt.Println(fileName + " deleted")
}
