package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura

func titulo(urls ...string) <-chan string {
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

func main() {
	t1 := titulo("https://www.canaldocampo.com.br/", "http://www.asmeduberaba.com.br/home")
	t2 := titulo("https://picpay.com/", "https://www.ifood.com.br/")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
