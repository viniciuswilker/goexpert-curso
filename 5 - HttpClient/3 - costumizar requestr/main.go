package main

import (
	"io"
	"net/http"
)

func main() {

	c := http.Client{}

	req, erro := http.NewRequest("GET", "http://google.com", nil)
	if erro != nil {
		panic(erro)
	}

	req.Header.Set("Accept", "application/json")
	resp, erro := c.Do(req)

	if erro != nil {
		panic(erro)
	}

	defer resp.Body.Close()

	body, erro := io.ReadAll(resp.Body)
	if erro != nil {
		panic(erro)
	}

	println(string(body))
}
