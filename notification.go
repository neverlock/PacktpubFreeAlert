package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const CLR_0 = "\x1b[30;1m"
const CLR_R = "\x1b[31;1m"
const CLR_G = "\x1b[32;1m"
const CLR_Y = "\x1b[33;1m"
const CLR_B = "\x1b[34;1m"
const CLR_M = "\x1b[35;1m"
const CLR_C = "\x1b[36;1m"
const CLR_W = "\x1b[37;1m"
const CLR_N = "\x1b[0m"

var color string
var server string
var msg string
var email string
var APIKEY = "Bearer [your APIKEY]"
var CHANNEL_TAG = "rusupport"

func getPackPubFree(Server string, MSG string) {
	urlStr := "https://api.pushbullet.com/v2/pushes"
	var alertString string
	if email == "someone@somedomain.com" {
		alertString = fmt.Sprintf("{\"type\": \"note\", \"title\": \"Notification from [%s]\", \"body\": \"%s !!!\",\"channel_tag\":\"%s\"}", Server, MSG, CHANNEL_TAG)
	} else {
		alertString = fmt.Sprintf("{\"type\": \"note\", \"title\": \"Notification from [%s]\", \"body\": \"%s !!!\",\"email\":\"%s\"}", Server, MSG, email)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(alertString))
	req.Header.Set("Authorization", APIKEY)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

func init() {
	flag.StringVar(&color, "c", "n", "default is n please use y for print color")
	flag.StringVar(&server, "server", "GPU1", "Server discription")
	flag.StringVar(&msg, "msg", "Your job dnoe", "Your msg for notification")
	flag.StringVar(&email, "email", "someone@somedomain.com", "Your pushbullet email or your email")
	flag.Parse()
}

func main() {

	if color != "y" {
		fmt.Println("+=====================================+")
		fmt.Println("|  Send notification................  |")
		fmt.Println("+=====================================+")
	} else {
		fmt.Println(CLR_G + "+=====================================+" + CLR_N)
		fmt.Println(CLR_G + "|  Send notification................  |" + CLR_N)
		fmt.Println(CLR_G + "+=====================================+" + CLR_N)
	}
	getPackPubFree(server, msg)
}
