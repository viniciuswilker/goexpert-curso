package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	c := http.Client{Timeout: time.Second * 2}
	resp, erro := c.Get("http://google.com")
	if erro != nil {
		panic(erro)
	}

	defer resp.Body.Close()

	corpoReq, erro := io.ReadAll(resp.Body)
	if erro != nil {
		panic(erro)
	}

	fmt.Println(string(corpoReq))
}
