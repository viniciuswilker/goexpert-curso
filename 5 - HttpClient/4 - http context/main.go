package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, erro := http.NewRequest("GET", "http://google.com", nil)
	if erro != nil {
		panic(erro)
	}

	resp, erro := http.DefaultClient.Do(req)
	if erro != nil {
		panic(erro)
	}
	defer resp.Body.Close()

	corpoReq, erro := io.ReadAll(resp.Body)
	if erro != nil {
		panic(erro)
	}

	println(string(corpoReq))

}
