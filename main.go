package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 0, "Select the port to run the web service")

	flag.Parse()

	if *port == 0 {
		flag.PrintDefaults()
	} else {

		fmt.Println("Print Headers and Body of HTTP requests. Port:", *port)

		http.HandleFunc("/", indexHandler)
		panic(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	fmt.Println("Header:")
	fmt.Println(r.Header)

	fmt.Println("Body:")
	fmt.Println(bodyString)
}
