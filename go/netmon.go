package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Connection struct {
	ContainerID   string  `json:"containerID"`
	ContainerName string  `json:"containerName"`
	ContainerPID  string  `json:"containerPID"`
	RequestIp     string  `json:"requestIP,omitempty"`
	Network       string  `json:"network"`
	OvsPortID     string  `json:"ovsPortID"`
	BandWidth     string  `json:"bandWidth,omitempty"`
	Delay         string  `json:"delay,omitempty"`
	RXTotal       uint64  `json:"rxKbytes"` // in KB
	TXTotal       uint64  `json:"txKbytes"` // in KB
	RXRate        float64 `json:"rxRate"`   // in Kb/s
	TXRate        float64 `json:"txRate"`   // in Kb/s

}

func main() {
	f, _ := os.Create("res.txt")
	defer f.Close()

	cid := os.Args[1]
	url := "http://127.0.0.1:8888/connection/" + cid
	for {
		resp, _ := http.Get(url)
		defer resp.Body.Close()

		var con Connection
		_ = json.NewDecoder(resp.Body).Decode(&con)

		fmt.Fprintf(f, "%s,%d,%d,%f,%f\n", con.Delay, con.RXTotal, con.TXTotal, con.RXRate, con.TXRate)
		time.Sleep(2 * time.Second)

	}
}
