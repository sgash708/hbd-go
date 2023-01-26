package main

import (
	"fmt"
	"time"

	age "github.com/sgash708/hbd-go"
)

func main() {
	birthDay := time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local)
	targetDay := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	myAge := age.CalAge(birthDay, targetDay)
	fmt.Println(myAge) // output: 1
}
