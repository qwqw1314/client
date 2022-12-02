package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for {
		reqBody := bytes.NewBufferString("Post plain text")
		resp, err := http.Post("loadbalancer.default.svc", "text/plain", reqBody)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(string(respBody))
		}
		time.Sleep(time.Second * 1)
	}
}
