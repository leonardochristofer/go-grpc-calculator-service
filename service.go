package main

import (
	"context"
	calculatorPb "go-grpc-calculator-service/pb/calculator"
	"math"
	"net"

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

func (s *server) FindPrimeNumber(ctx context.Context, CalculatorRequest *calculatorPb.CalculatorRequest) (*calculatorPb.CalculatorReturn, error) {
	n := int(CalculatorRequest.GetN())

	if n <= 0 {
		n = 1
	}

	counter := 0

	currNum := 1

	for counter < n {
		currNum++
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(currNum))); i++ {
			if currNum%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			counter++
		}
	}

	return &calculatorPb.CalculatorReturn{Result: int64(currNum)}, nil
}

func PrimeNumber(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func ReverseNumber(n int) int {
	reverse := 0
	for {
		reverse = (reverse * 10) + (n % 10)
		n /= 10
		if n == 0 {
			break
		}
	}
	return reverse
}

func PalindromeNumber(n int) bool {
	return (n == ReverseNumber(n))
}

func (s *server) FindPrimePalindromeNumber(ctx context.Context, CalculatorRequest *calculatorPb.CalculatorRequest) (*calculatorPb.CalculatorReturn, error) {
	n := int(CalculatorRequest.GetN())

	if n <= 0 {
		n = 1
	}

	count := 0

	primePalindrome := 0

	for i := 0; count < n; i++ {
		if PrimeNumber(i) && PalindromeNumber(i) {
			primePalindrome = i
			count++
		}
	}

	return &calculatorPb.CalculatorReturn{Result: int64(primePalindrome)}, nil
}
