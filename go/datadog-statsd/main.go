package main

import (
	"log"

	"github.com/DataDog/datadog-go/v5/statsd"
)

func main() {
	client, err := statsd.New("127.0.0.1:8125",
		statsd.WithTags([]string{"env:dev", "service:myservice"}),
	)
	if err != nil {
		log.Fatal(err)
	}

	client.Histogram("my.metrics", 21, []string{"tag1", "tag2:value"}, 1)
	client.Close()
}
