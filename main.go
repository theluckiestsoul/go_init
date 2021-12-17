package main

import (
	"fmt"
	"time"
)

var weekday string

func init() {
	weekday = time.Now().Weekday().String()
}

func init() {
	fmt.Println("First init")
}

func init() {
	fmt.Println("Second init")
}

func init() {
	fmt.Println("Third init")
}

func init() {
	fmt.Println("Fourth init")
}

func main() {
	fmt.Printf("Today is %s", weekday)
}
