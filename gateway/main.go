package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	str := "hello world, I am http body"

	for i := 0; i < 5; i++ {
		req, err := http.NewRequest("GET", "http://127.0.0.1/", bytes.NewReader([]byte(str)))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("If-None-Match", `W/"wyzzy"`)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s \n", resp)
	}
}
