package main

import (
    f "fmt"
    "net/http"
)

func main() {

    f.Println("running http server at 8080 port")

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        res.Write([]byte("Hello World"))
    })

    http.ListenAndServe(":8080", nil)
}
