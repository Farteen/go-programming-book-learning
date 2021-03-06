package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			fmt.Println(url)
			url = "http://" + url
		}
		resp, err := http.Get(url)
		fmt.Println(resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err);
			os.Exit(1)
		}
		result, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", result)
	}
}