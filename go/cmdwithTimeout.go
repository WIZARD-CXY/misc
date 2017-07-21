package main

import (
   "fmt"
   "os/exec"
   "time"
   "io/ioutil"
)

func main(){
   cmd := exec.Command("find", "/")
   stdoutp, err := cmd.StdoutPipe()
   
   if err != nil {
     return
   }


   stderrp, err := cmd.StderrPipe()

   if err != nil {
     return 
   }

    if err := cmd.Start(); err != nil {
       fmt.Println("ahah")
       return
    }
   
    timer := time.AfterFunc(5*time.Second, func(){
         fmt.Println("killing cmd due to timeout")
         cmd.Process.Kill()
    })

    stdoutb, souterr := ioutil.ReadAll(stdoutp)

    if souterr != nil {
        fmt.Println("failed to read from stdout")
    }


    stderrb, _ := ioutil.ReadAll(stderrp)
   
    err = cmd.Wait()
    timer.Stop()

    fmt.Println(string(stdoutb),"++++", string(stderrb),err)
}
