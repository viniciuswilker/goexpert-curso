package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second * 2}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "wesley"}`))

	resp, erro := c.Post("http://google.com", "application/json", jsonVar)
	if erro != nil {
		panic(erro)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
