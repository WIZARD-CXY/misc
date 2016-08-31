package main
// ref https://serholiu.com/go-http-client-keepalive
import (
    "net"
    "time"
    "net/http"
    "fmt"
    "io/ioutil"
)

func PrintLocalDial(network, addr string) (net.Conn, error){
    dial := net.Dialer{
        Timeout: 1*time.Second,
        KeepAlive: 30*time.Second,
    }

    conn, err := dial.Dial(network,addr)
    if err!=nil{
        return conn, err
    }

    fmt.Println("connect done, use", conn.LocalAddr().String())

    return conn, err
}

const URL1 ="http://www.baidu.com"
const URL2 ="https://www.baidu.com/search/error.html"
func main(){
    client := &http.Client{
        Transport: &http.Transport{
            Dial: PrintLocalDial,
            MaxIdleConnsPerHost:3,
        },
    }

    for {
        go doGet(client, URL1,1)
        go doGet(client, URL1,2)
        go doGet(client, URL1,3)
     
        time.Sleep(2*time.Second)
    }
}

func doGet(client *http.Client, url string, id int){
    resp, err := client.Get(url)

    if err!=nil{
        fmt.Println(err)
        return
    }

    buf, err := ioutil.ReadAll(resp.Body)

    fmt.Printf("%d: %s -- %v\n",id, string(buf[:10]), err)

    resp.Body.Close()
}