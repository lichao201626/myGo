package main

import (
    "fmt"
    "net/http"
    "log"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello there!\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

/* func main(){
    http.HandleFunc("/", myHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
