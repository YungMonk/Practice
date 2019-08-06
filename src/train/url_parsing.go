package train

import (
	"fmt"
	"net/url"
)

// URLParsing URL解析
func URLParsing() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(url.ParseQuery(u.RawQuery))
}
