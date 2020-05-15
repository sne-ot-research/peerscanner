package scanner

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func CosScan(ips []string) map[string][][]string {
	var wg sync.WaitGroup
	resChan := make(chan map[string][][]string)

	// for each ip
	// add to wait group
	// find cors setting in a go routine
	for _, ip := range ips {
		wg.Add(1)
		fmt.Println("scanning:", ip)
		go findCorsSetting(&wg, resChan, ip)
	}

	// wait till all goroutines are done then close channel
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// read off the closed channel
	result := make(map[string][][]string)
	for r := range resChan {
		for k, v := range r {
			result[k] = v
		}
	}

	return result
}

func findCorsSetting(wg *sync.WaitGroup, c chan<- map[string][][]string, ip string) {
	defer wg.Done()
	// recover from the error
	defer func() {
		if e := recover(); e != nil {
			log.Fatal(e)
		}
	}()

	client := http.Client{Timeout: time.Duration(30 * time.Second)}
	request, err := http.NewRequest("OPTIONS", fmt.Sprintf("http://%s:%d/api/v0/add", ip, 5001), nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Origin", "*")

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	c <- map[string][][]string{
		ip: {
			resp.Header["Access-Control-Allow-Origin"],
			resp.Header["Headers.Access-Control-Allow-Credentials"],
		},
	}
}