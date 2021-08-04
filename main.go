package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        name, err := os.Hostname()
        if err != nil {
            panic(err)
        }
        fmt.Println("Serving request!")
        fmt.Fprintf(w, "Hello from %v", name)
    })

    log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

