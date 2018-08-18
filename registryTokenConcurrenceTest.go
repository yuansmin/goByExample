/* test docker registry token server's concurrency
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	concurrency = 10
	total       = 100
)

var url = "http://<host>/registry/token?scope=repository%3Alibrary%2Fbusybox%3Apush%2Cpull&service=<registry-service>"
var headers = map[string]string{"...": "..."}

func main() {
	fmt.Println("to finish")
	start := time.Now()
	ch := make(chan *TestData, 4)
	limitCh := make(chan int, 4)
	go feed(ch, limitCh)

	var successCount int
	var sum float64
	for i := 0; i < total; i++ {
		d := <-ch
		<-limitCh
		sum += d.seconds
		fmt.Println(d)
		if d.ok {
			successCount++
		}
	}
	cost := time.Since(start).Seconds()
	average := sum / float64(total)
	fmt.Printf("\nsuccessCount: %d\ncost: %.2f    average: %.2f\n", successCount, cost, average)
}

type TestData struct {
	seconds    float64 // seconds that the request cost
	ok         bool    // is request success
	statusCode int     // response status code
	message    string  // error message
}

func (d *TestData) String() string {
	return fmt.Sprintf("ok: %v\nstatusCode:%d\ncost: %.2f\nmessage: %s\n",
		d.ok, d.statusCode, d.seconds, d.message)
}

func feed(ch chan<- *TestData, limitCh chan<- int) {
	for i := 0; i < total; i++ {
		limitCh <- 1
		go getToken(ch)
	}
}

func getToken(ch chan<- *TestData) {
	start := time.Now()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		secs := time.Since(start).Seconds()
		d := &TestData{secs, false, -1, fmt.Sprint(err)}
		ch <- d
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		secs := time.Since(start).Seconds()
		d := &TestData{secs, false, resp.StatusCode, fmt.Sprint(err)}
		ch <- d
		return
	}
	secs := time.Since(start).Seconds()
	ok := resp.Body.StatusCode == 200
	d := &TestData{secs, ok, resp.StatusCode, string(body)}
	ch <- d
}
