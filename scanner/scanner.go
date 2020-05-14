package scanner

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CosScan(ips []string) map[string][][]string {

	res := make(map[string][][]string)

	for _, ip := range ips {

		fmt.Println("scanning:", ip)

		client := http.Client{Timeout: time.Duration(30 * time.Second)}
		request, err := http.NewRequest("OPTIONS", fmt.Sprintf("http://%s:%d/api/v0/add", ip, 5001), nil)
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Set("Origin", "*")

		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
			continue
		}

		res[ip] = [][]string{
			resp.Header["Access-Control-Allow-Origin"],
			resp.Header["Headers.Access-Control-Allow-Credentials"],
		}
	}

	return res
}
