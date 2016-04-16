package main

// qiniu list bucket
import (
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

func main() {

	AK := "xxxx"
	SK := "xxxx"
	bucketName := "xxx"
	kodo.SetMac(AK, SK)

	zone := 0
	c := kodo.New(zone, nil)
	bucket := c.Bucket(bucketName)

	ctx := context.Background()

	res, _, _, _ := bucket.List(ctx, "", "", "", 10)

	for _, file := range res {
		fmt.Println(file)
	}

}
