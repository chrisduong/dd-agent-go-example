package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!!")
}

func main() {
	ddSock := os.Getenv("DD_DOGSTATSD_SOCKET")

	// ddHost := os.Getenv("DATADOG_HOST")
	// if len(ddHost) == 0 {
	// 	ddHost = "127.0.0.1"
	// }

	// statsd, err := statsd.New(fmt.Sprintf("%s:8125", ddHost),
	// 	statsd.WithTags([]string{"env:test", "service:go-example"}),
	// )
	statsd, err := statsd.New(fmt.Sprintf("unix://%s", ddSock),
		statsd.WithTags([]string{"env:test", "service:go-example"}),
	)
	if err != nil {
		log.Fatal(err)
	}

	statsd.Gauge("mygo.metrics", 21, []string{"tag1", "tag2:value"}, 1)
	// statsd.Close()

	var i float64
	go func() {
		for {
			i++
			statsd.Gauge("example_metric.gauge", i, []string{"environment:test"}, 1)
			// log.Println("Sent!")
			time.Sleep(10 * time.Second)
		}
	}()

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
