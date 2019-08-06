package train

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Host:", r.URL.Host)
	fmt.Println("URL:", r.URL)
	for key, val := range r.Form {
		fmt.Println("key : ", key)
		fmt.Println("val : ", val)
	}
	fmt.Fprintf(w, "Hello World!")
}

// MyHTTP 搭建 http
func MyHTTP() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listen And Server :", err)
	}
}
