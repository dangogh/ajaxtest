package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/time", timeData)
	http.HandleFunc("/base.css", getCSS)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start.")
	}
}

func timeData(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprint(w, t.String() )
}

func HomePage(writer http.ResponseWriter, request *http.Request) {
	inFile, err := os.Open("html/homepage.html")
	if err != nil {
		log.Fatal(err)
	}
	page := ""
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		page += scanner.Text() + "\n"
	}
	err = scanner.Err()
	if err == io.EOF {
		log.Print("EOF Reached")
	}

	fmt.Fprint(writer, page)
}

func getCSS(writer http.ResponseWriter, request *http.Request) {
	inFile, err := os.Open("css/base.css")
	if err != nil {
		log.Fatal(err)
	}
	page := ""
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		page += scanner.Text() + "\n"
	}
	err = scanner.Err()
	if err == io.EOF {
		log.Print("EOF Reached")
	}

	fmt.Fprint(writer, page)
}
