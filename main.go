package main

import (
    "fmt"
    "net/http"
    "sort"
)

func handler(w http.ResponseWriter, r *http.Request) {

    var keys []string
    for k := range r.Header {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    fmt.Fprintln(w, "<b>Request Headers:</b></br>", r.URL.Path[1:])
    for _, k := range keys {
        fmt.Fprintln(w, k, ":", r.Header[k], "</br>", r.URL.Path[1:])
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}