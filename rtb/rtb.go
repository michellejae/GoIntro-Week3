package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 10 * time.Millisecond

	bid := bidOn("http://a.b", timeout)
	fmt.Println(bid)

	bid = bidOn("http://ardanlabs.com", timeout)
	fmt.Println(bid)
}

// RTB: Real Time Bidding

func bidOn(url string, timeout time.Duration) Bid {
	ch := make(chan Bid, 1) // buffered channel to avoid goroutine leak
	go func() {
		ch <- bestBid(url)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case bid := <-ch: //algo finished in time
		return bid
	case <-ctx.Done(): // timeout
		return defaultBid
	}
}

var defaultBid = Bid{
	Price: 0.05,
	URL:   "https://adsRus.com/ad/default",
}

type Bid struct {
	Price float64
	URL   string
}

// Code written by algo team
func bestBid(url string) Bid {
	// simulalte algo work
	time.Sleep(time.Duration(len(url)) * time.Millisecond)
	return Bid{
		Price: 0.03,
		URL:   "https://adsRus.com/ad/7",
	}
}
