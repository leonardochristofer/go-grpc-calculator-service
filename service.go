package main

import (
	"context"
	calculatorPb "go-grpc-calculator-service/pb/calculator"
	"math"
	"net"
	"sort"
	"strconv"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	calculatorPb.RegisterCalculatorServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func IsPrime(num int64) bool {
	if num == 2 || num == 1 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	lim := int64(math.Sqrt(float64(num)))
	var i int64
	for i = 3; i <= lim; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func IsPalindrome(num int64) bool {
	str := strconv.FormatInt(num, 10)
	strlen := len(str)
	for i := 0; i < strlen/2; i++ {
		if str[i] != str[strlen-1-i] {
			return false
		}
	}
	return true
}

type Int64s []int64

func (is Int64s) Len() int {
	return len(is)
}
func (is Int64s) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is Int64s) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (s *server) FindPrimeNumber(ctx context.Context, CalculatorRequest *calculatorPb.CalculatorRequest) (*calculatorPb.CalculatorReturn, error) {
	n := int(CalculatorRequest.GetN())
	// convert n to offset
	n -= 1
	// goroutine stuff
	var wg sync.WaitGroup
	var candidates = make(chan int64)
	var done = make(chan struct{})
	var candidatesResults = make(chan int64)
	var results = make([]int64, 0, n)
	// checking workers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range candidates {
				if IsPrime(c) {
					candidatesResults <- c
				}
			}
		}()
	}
	// send candidates
	go func() {
		// send candidates out
		for i := int64(2); ; i++ {
			select {
			case _, ok := <-done:
				if !ok {
					close(candidates)
					return
				}
			case candidates <- i:
				continue
			}
		}
	}()
	// collect results
	go func() {
		for c := range candidatesResults {
			results = append(results, c)
			if len(results) == n+1 {
				// we have enough, just need to make sure we got all of them
				close(done)
			}
		}
	}()
	wg.Wait()
	// sort array ascending
	sort.Sort(Int64s(results))
	// return results based on index
	return &calculatorPb.CalculatorReturn{Result: int64(results[n])}, nil
}

func (s *server) FindPrimePalindromeNumber(ctx context.Context, CalculatorRequest *calculatorPb.CalculatorRequest) (*calculatorPb.CalculatorReturn, error) {
	n := int(CalculatorRequest.GetN())
	// convert n to offset
	n -= 1
	// goroutine stuff
	var wg sync.WaitGroup
	var candidates = make(chan int64)
	var done = make(chan struct{})
	var candidatesResults = make(chan int64)
	var results = make([]int64, 0, n)
	// checking workers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range candidates {
				if IsPalindrome(c) && IsPrime(c) {
					candidatesResults <- c
				}
			}
		}()
	}
	// send candidates
	go func() {
		// send candidates out
		for i := int64(2); ; i++ {
			select {
			case _, ok := <-done:
				if !ok {
					close(candidates)
					return
				}
			case candidates <- i:
				continue
			}
		}
	}()
	// collect results
	go func() {
		for c := range candidatesResults {
			results = append(results, c)
			if len(results) == n+1 {
				// we have enough, just need to make sure we got all of them
				close(done)
			}
		}
	}()
	wg.Wait()
	// sort array ascending
	sort.Sort(Int64s(results))
	// return results based on index
	return &calculatorPb.CalculatorReturn{Result: int64(results[n])}, nil
}
