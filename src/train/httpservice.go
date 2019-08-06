package train

import (
	"fmt"
	"log"
	"net/http"
)

type selfHTTP struct {
	Data interface{}
	Resp http.ResponseWriter
	Requ *http.Request
}

func (h *selfHTTP) ServeHTTP(resp http.ResponseWriter, requ *http.Request) {
	h.Resp = resp
	h.Requ = requ
	h.Data = "Use define serverhttp"
	resp.Write([]byte(h.Data.(string)))
}

// CreateHTTPService 创建HTTP服务
func CreateHTTPService() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		fmt.Println("PATH: ", req.URL.Path)
		fmt.Println("SCHEME: ", req.URL.Scheme)
		fmt.Println("METHOD: ", req.Method)
		fmt.Println()

		fmt.Fprintf(res, "<h1>Index Page</h1>")

		p := &selfHTTP{}
		p.ServeHTTP(res, req)
	})

	log.Fatal(http.ListenAndServe(":80", &selfHTTP{}))
}
