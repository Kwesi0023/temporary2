package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(

	http.HandlerFunc("/static", serveStatic)
	http.HandlerFunc("/", serveDynamic)
	http.ListenAndServe (Port,nil)
)




const(
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request){
	response:= "The time is currently " + time.Now().String()
	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}