package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://example.com",
		"http://amazon.de",
	}

	channel := make(chan string) // creates empty channel. required to use make()

	// for _, link := range links {
	// 	go checkLink(link, channel)
	// }

	/*	simple approach: */
	// fmt.Println(<-channel) // consume the content of channel
	// fmt.Println(<-channel)
	// third execution is omitted, as it's not waited for

	/* more "sophisticated" approach, if you want to wait on each value */
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-channel)
	// }

	// attempt an immediate refetch forever over
	// for {
	// 	go checkLink2(<-channel, channel)
	// }
	// === totally equal to ===
	// for link := range channel {
	// 	go checkLink2(link, channel)
	// }

	for _, link := range links {
		go checkLink2(link, channel)
	}

	// call infinite and wait after each execution
	for l := range channel {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink2(link, channel)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode > 299 {
		fmt.Println(link, "is down")
		c <- "it's down"
		return
	}
	fmt.Println(link, "is up")
	c <- "it's up"
}

func checkLink2(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode > 299 {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
