package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// resp, err := http.Get("http://localhost:5000")
	// resp, err := http.Get("http://web01.meluskyc.org:8080")
	resp, err := http.Get("http://d153fb5f.ngrok.io")

	if err != nil {
		log.Print("send error")
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print("read error")
		log.Fatal(err)
	}
	fmt.Println(os.Stdout, string(data))
}
