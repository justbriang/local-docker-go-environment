package main

import (
    "fmt"  
    // "html"
    "log"
    "net/http"
)

func main() {
    // log.Printf("hello World")


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        _ ,err := fmt.Fprintf(w, "Hello world 5")
        if err != nil {
            fmt.Println(err)
            // return nil, err
        }
        
    })

    // http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "Hi")
    // })

    log.Fatal(http.ListenAndServe(":80", nil))

}