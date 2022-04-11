package main

import (
	"fmt"
	"time"
)

const layout = "01-02-2006"

func getAge(dob string) int {
	birthday, _ := time.Parse(layout, dob)
	now := time.Now().Format(layout)
	today, _ := time.Parse(layout, now)

	ty, tm, td := today.Date()
	by, bm, bd := birthday.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	birthday = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthday) {
		return 0
	}
	return ty - by

}

func main() {
	var dob string
	var age int
	fmt.Scan(&dob)
	age = getAge(dob)
	fmt.Println(age)
}
