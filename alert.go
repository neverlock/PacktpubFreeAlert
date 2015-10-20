package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
	"gopkg.in/gcfg.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type Config struct {
	Pushbullet struct {
		APIKEY string
		Cron   string
	}
}

var cfg Config

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
var FreeBook string
var Updater int

func getPackPubFree() {
	url := "https://www.packtpub.com/packt/offers/free-learning"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	book := strings.TrimSpace(doc.Find(".dotd-title h2").Text())

	if FreeBook != book {
		FreeBook = book
		Updater = 1
	} else {
		Updater = 0
	}

	if Updater == 1 {
		log.Println(book)
		log.Println("New Free Book update")
		urlStr := "https://api.pushbullet.com/v2/pushes"
		alertString := fmt.Sprintf("{\"type\": \"link\", \"title\": \"Book Update\", \"body\": \"%s !!!\",\"url\":\"%s\"}", book, url)
		client := &http.Client{}
		req, _ := http.NewRequest("POST", urlStr, strings.NewReader(alertString))
		req.Header.Set("Authorization", cfg.Pushbullet.APIKEY)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := client.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(body))
	} else {
		log.Println("No Free Book update")
	}
}

func init() {
	flag.StringVar(&color, "c", "n", "default is n please use y for print color")
	flag.Parse()
}

func main() {

	err := gcfg.ReadFileInto(&cfg, "configAlert.gcfg")
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	if color != "y" {
		fmt.Println("+=====================================+")
		fmt.Println("|  Checking for page update.........  |")
		fmt.Println("+=====================================+")
		fmt.Println("Use CTRL+C to Exit")
	} else {
		fmt.Println(CLR_G + "+=====================================+" + CLR_N)
		fmt.Println(CLR_G + "|  Checking for page update.........  |" + CLR_N)
		fmt.Println(CLR_G + "+=====================================+" + CLR_N)
		fmt.Println(CLR_R + "Use CTRL+C to Exit" + CLR_N)
	}
	c := cron.New()
	c.AddFunc(cfg.Pushbullet.Cron, func() { getPackPubFree() })
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}
