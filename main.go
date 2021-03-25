package main

import (
	"fairBilling/service"
	"fmt"
	"os"
)

func main() {
	//read the file path from std input
	file := os.Args[1]
	newSession := service.NewFairBillingService(file)
	result := newSession.CalculateSession()
	//print the result
	for name, session := range *result {
		fmt.Printf("%s %d %v \n", name, session.SessionCount, session.Duration)
	}
}
