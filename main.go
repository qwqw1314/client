package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var ip string
	file, err := os.Open("/etc/hostname")
	if err != nil {
		panic(err)
	}
	hostname := bufio.NewScanner(file)
	hostname.Scan()
	subtext := hostname.Text()
	file.Close()

	file, err = os.Open("/etc/hosts")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, subtext) {
			ip = strings.Split(text, " ")[0]
		}
	}

	file.Close()

	for {
		reqBody := bytes.NewBufferString(ip)
		resp, err := http.Post("http://loadbalancer.default.svc", "text/plain", reqBody)
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
