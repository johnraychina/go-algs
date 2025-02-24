package main

import (
	"context"
	"fmt"
	"time"
)
import "golang.org/x/time/rate"

// detect max io rate
// avg time
// penalize

func main() {

	//metricsFile, err := os.Create("metrics.txt")
	//check(err)
	//defer metricsFile.Close()
	//metricsWriter := bufio.NewWriter(metricsFile)

	ctx := context.Background()
	limiter := rate.NewLimiter(100, 100)
	windowSize := 10 // small window size can be more sensitive to time cost changes
	windowCount := 100
	var costs []time.Duration
	//var penaltys []time.Duration

	// window1
	var lastCost time.Duration
	for i := 0; i < windowSize; i++ {
		lastCost += doIO(ctx, ioFunction, limiter)
	}
	costs = append(costs, lastCost)
	fmt.Println(lastCost)

	for n := 0; n < windowCount; n++ {

		// slide window
		curCost := time.Duration(0)
		for i := 0; i < windowSize; i++ {
			curCost += doIO(ctx, ioFunction, limiter)
		}
		// collect data to draw  time cost for each window
		costs = append(costs, curCost)
		fmt.Println(curCost)

		// compare and penalize
		slow := curCost - lastCost
		penalty := time.Duration(0)
		if slow > 0 {
			penalty = slow * 2
			//penaltys = append(penaltys, penalty)
			fmt.Println("penalty", penalty)
		}
		time.Sleep(penalty) // time passed and limiter is generating more token(io resource is recovery from high pressure)

		lastCost = curCost
	}

	//for _, cost := range costs {
	//	fmt.Println(cost)
	//}

	// todo assume io limiter would change over time, can we use kalman filter to follow the io limiter?
}

func check(err error) {
	panic(err)
}

func ioFunction() {

}

func doIO(ctx context.Context, f func(), limiter *rate.Limiter) time.Duration {
	start := time.Now()
	// mock io function and its rate
	f()
	time.Sleep(10 * time.Millisecond)
	err := limiter.Wait(ctx)
	if err != nil {
		panic(err)
	}
	return time.Since(start)
}
