// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	hello := []byte("Hello World!!!")
// 	_, err := w.Write(hello)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/hello", helloHandler)
// 	fmt.Println("Server Start Up........")
// 	log.Fatal(http.ListenAndServe("localhost:8080", nil))
// }


package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}