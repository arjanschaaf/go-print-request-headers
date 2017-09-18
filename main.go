package main

import (
    "fmt"
    "net/http"
    "sort"
)

func handler(w http.ResponseWriter, r *http.Request) {

    // adding debug header to test (strong/weak) ETags in combination with NGINX
    w.Header().Set("ETag", "HelloWorld")

    var requestKeys []string
    for k := range r.Header {
        requestKeys = append(requestKeys, k)
    }
    sort.Strings(requestKeys)

    var responseKeys []string
    for k := range w.Header() {
        responseKeys = append(responseKeys, k)
    }
    sort.Strings(responseKeys)

    fmt.Fprintln(w, "<b>request.RequestURI:</b>", r.RequestURI, "</br>")
    fmt.Fprintln(w, "<b>request.RemoteAddr:</b>", r.RemoteAddr, "</br>")
    fmt.Fprintln(w, "<b>request.TLS:</b>", r.TLS, "</br>")


    fmt.Fprintln(w, "<b>Request Headers:</b></br>")
    for _, k := range requestKeys {
        fmt.Fprintln(w, k, ":", r.Header[k], "</br>")
    }

    /*
    fmt.Fprintln(w, "<b>Response Headers:</b></br>")
    for _, k := range responseKeys {
        fmt.Fprintln(w, k, ":", k, "</br>")
    }
    */
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}