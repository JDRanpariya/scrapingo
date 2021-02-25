package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	URL := "https://vixen.com"

	client := &http.Client{}

	res, err := client.Get(URL)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.StatusCode)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
