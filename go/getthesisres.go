package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/subosito/twilio"
)

const senderNumber = "+13852122788"
const recvNumber = "+*"

var resUrl = "http://101.200.157.98/student"
var loginUrl = "http://101.200.157.98/login"

func sendSMS(res string) {
	var AccountSID, Auth_Token string

	if os.Getenv("SID") == "" {
		AccountSID = "AC30de90969d9b2fa749a20c52813f1c6e"
	} else {
		AccountSID = os.Getenv("SID")
	}

	if os.Getenv("TOKEN") == "" {
		Auth_Token = "*"
	} else {
		Auth_Token = os.Getenv("TOKEN")
	}

	client := twilio.NewClient(AccountSID, Auth_Token, nil)
	message := fmt.Sprint("message from ", res)

	params := twilio.MessageParams{
		Body: message,
	}

	s, resp, err := client.Messages.Send(senderNumber, recvNumber, params)

	if err != nil {
		log.Fatal(s, resp, err)
	}
}

func getRes() string {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	form := url.Values{}
	form.Add("username", "*")
	form.Add("password", "*")

	resp, err := client.PostForm(loginUrl, form)
	//resp.Body.Close()

	//resp, err = client.Get(resUrl)

	if err != nil {
		log.Fatal("get res fail")
	}

	htmlData, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(htmlData)

}
func main() {
	for i := 0; ; i++ {
		log.Println(i)
		resp := getRes()

		if !strings.Contains(resp, "尚无结果") {
			sendSMS("wiz")
			return
		}
		time.Sleep(5 * time.Minute)

	}

}
