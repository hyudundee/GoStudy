package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://vcfgdsfgfssdgffsfrefred.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	/*
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c) // one extra print will hang up the program
		// becase message coming from c is a blocking call
	*/
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Blocking call
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, " is up")
	c <- "Yep it is up"
}
