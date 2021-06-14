package main

import (
   "fmt"
   "os"
   "path/filepath"
   "io/ioutil"
)

func main() {
     dir := filepath.Dir(filepath.Dir(os.Args[0]))

     fmt.Println(dir)

    files, err := ioutil.ReadDir(dir)

    if err != nil {
        fmt.Println(err)
        return
    }
   
   for _, file := range files {
       fmt.Println(dir+"/"+file.Name())
   }

}
