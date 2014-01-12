package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/time", timeData)
	http.HandleFunc("/css/base.css", getCSS)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start.")
	}
} //main

func timeData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().String() )
}

func HomePage(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "html/homepage.html")
}

func getCSS(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "css/base.css")
}
