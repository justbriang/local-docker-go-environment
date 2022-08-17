// package main

// import (
// 	"log"
// 	"net/http"
// 	// "/app/cmd/web/routes.go"
// )
// func main(){

// 	rt:=routes()
	
// 	log.Println("Listening on port 8080")
// 	_ = http.ListenAndServe(":80",rt)

// }

// package main

// import (
//     "fmt"
// 	"log"
//     "net/http"
// )

// func main() {
//     http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//     })
// 	log.Println("Listening on port 8080")
//     http.ListenAndServe(":80", nil)
// }

package main

import (
    // "fmt"
	"log"
    "net/http"
)
 func main(){
	//     http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    // })
    routes:=routes()

	log.Println("Listening on port 8080")
    http.ListenAndServe(":80", routes)
 }