package main

import (
	"fmt"
	calculatorPb "go-grpc-calculator-service/pb/calculator"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := calculatorPb.NewCalculatorServiceClient(conn)

	g := gin.Default()

	g.GET("/api/prime/:n", func(ctx *gin.Context) {
		n, err := strconv.ParseUint(ctx.Param("n"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
			return
		}
		req := &calculatorPb.CalculatorRequest{N: int64(n)}

		if CalculatorReturn, err := client.FindPrimeNumber(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(CalculatorReturn.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/api/prime/palindrome/:n", func(ctx *gin.Context) {
		n, err := strconv.ParseUint(ctx.Param("n"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
			return
		}

		req := &calculatorPb.CalculatorRequest{N: int64(n)}

		if CalculatorReturn, err := client.FindPrimePalindromeNumber(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(CalculatorReturn.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":9100"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
