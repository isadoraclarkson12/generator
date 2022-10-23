package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura

func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
