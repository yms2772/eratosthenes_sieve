package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	num := flag.Int("max", 0, "최대 숫자")
	flag.Parse()

	if *num < 2 {
		fmt.Println("'max'는 1보다 커야합니다")

		return
	}

	startTime := time.Now()

	var prime []int
	isPrime := make([]bool, *num+1)

	for i := 2; i <= *num; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= *num; i++ {
		if isPrime[i] {
			for j := i * i; j <= *num; j += i {
				isPrime[j] = false
			}
		}
	}

	for primeNum, b := range isPrime {
		if b {
			prime = append(prime, primeNum)
		}
	}

	fmt.Printf("소요시간: %vms", time.Since(startTime).Milliseconds())
	#fmt.Println(prime)
}
