package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type result struct {
	i   int
	avg float64
	err error
}

func main() {
	// declaring our context to kno whow long to go if something gets stuck somewhere. (could have used a select statement)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	out, err := pingTimes(ctx, []string{"google.com", "ibm.com", "apple.com"}, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

}

func parsePing(out string) (float64, error) {
	// get bottom line of ping return that starts with 'round trip'
	bottomLine := strings.SplitAfter(out, "round-trip")
	// split the bottom line to grab by '/'
	splits := strings.Split(bottomLine[1], "/")
	// grab the time that we know is the average
	value := splits[4]
	// converst it to a numero
	average, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return average, nil
}

func pingTimes(ctx context.Context, hosts []string, count int) ([]time.Duration, error) {
	// passing context through here as well cause i want to make sure it cancels if our actually pings do not return
	c := "-c"
	ch := make(chan result)
	// loop through hosts
	for i, host := range hosts {
		// call a go routine for each loop through our hosts
		go func(ctx context.Context, i int, host string, count int) {
			// ping the hosts
			out, err := exec.CommandContext(ctx, "ping", host, c, fmt.Sprintf("%d", count)).Output()
			if err != nil {
				log.Fatal(err)
			}
			// convert return to string
			time := string(out)
			// parse the return to get back the avg time
			avg, err := parsePing(time)
			if err != nil {
				log.Fatal(err)
			}
			// send the index, avg, and any error as type struct to channel (this is cause we created channel as struct since it can only be once type but struct can have multiple types)
			ch <- result{i, avg, err}
		}(ctx, i, host, count)
	}

	out := make([]time.Duration, len(hosts))
	// have to loop through hosts again to recevive in channel (if we do not do this we will only get data from one go routine cause next one cannot start until channel reads)
	for range hosts {
		val := <-ch // receive
		if val.err != nil {
			return nil, val.err
		}
		// save data to new slice we created
		out[val.i] = time.Duration(val.avg) * time.Millisecond
	}
	return out, nil
}
