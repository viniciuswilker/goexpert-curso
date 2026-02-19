package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Rodando CLIENT")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	res, erro := http.DefaultClient.Do(req)

	if erro != nil {
		panic(erro)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, req.Body)

}
