package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	fmt.Println(linspace(1, 7, 10))

	values := []float64{2, 1, 3}
	m, err := median(values)
	if err != nil {
		log.Fatal(err) // short version of log.Printf(eff) + os.Exit(1)
	}
	fmt.Println("median", values, "->", m)

	values = append(values, 4)
	m, err = median(values)
	if err != nil {
		log.Fatal(err) // short version of log.Printf(eff) + os.Exit(1)
	}
	fmt.Println("median", values, "->", m)
}

/*
- sort values
- if len(values) is odd - return middle value
- otherwise return average of two middle values
*/
func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("medium of empty slice")
	}

	// doesn't change values when you sort
	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	mid := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[mid], nil
	}
	// able to divide by 2 here without converting cause we just use value that's
	// type is determined when i use it VS if we were like n = 2 then divide by n
	return (nums[mid-1] + nums[mid]) / 2, nil
}

// linspace (0, 1, 10) -> [0, 0.1, .... 1]
func linspace(start, stop float64, count int) []float64 {
	step := (stop - start) / float64(count-1)
	// var vec []float64 // option one
	//vec := make([]float64, count) // option 2
	vec := make([]float64, 0, count) // option 3
	// currCap := cap(vec) // just used to test capacity
	for i := 0; i < count; i++ {
		val := start + step*float64(i)
		vec = append(vec, val) // option 1 and option 3
		//	vec[i] = val option 2
		// if c := cap(vec); c != currCap {
		// 	fmt.Printf("allocation: %d -> %d\n", currCap, c)
		// 	currCap = c
		// }
	}
	return vec
}
