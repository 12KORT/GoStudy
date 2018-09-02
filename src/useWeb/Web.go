package useWeb

import (
	"fmt"
	"log"
	"net/http"
)

func SimpleHttpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}

	fmt.Fprintf(w, "hello, it's test web")
}

func WebListen() {
	http.HandleFunc("/", SimpleHttpHandler)
	fmt.Print("begin web listen")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}
