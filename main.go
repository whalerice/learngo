package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main()  {
	var results = map[string]string{}
	urls := []string {
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
		
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)

		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}
	
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string)  error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		return errRequestFailed
	}
	return nil
}