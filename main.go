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

    fmt.Fprintln(w, "<b>request.RequestURI:</b>", r.RequestURI, "</br>")
    fmt.Fprintln(w, "<b>request.RemoteAddr:</b>", r.RemoteAddr, "</br>")
    fmt.Fprintln(w, "<b>request.TLS:</b>", r.TLS, "</br>")



    fmt.Fprintln(w, "<b>Request Headers:</b></br>")
    for _, k := range keys {
        fmt.Fprintln(w, k, ":", r.Header[k], "</br>")
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}