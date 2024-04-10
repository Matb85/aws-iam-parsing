package main

import (
	"fmt"
	"matb85/remitly-home-assignment/method"
	"os"
)

func main() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	result := method.VerifyPolicyJSON(data)
	fmt.Println((result))
}
