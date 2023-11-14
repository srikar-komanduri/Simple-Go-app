package main

import (
	"fmt"
	"time"
)

const layout = "01-02-2008"

func getAge(dob string) int {
	birthday, _ := time.Parse(layout, dob)
	// asdasd
	now := time.Now().Format(layout)
	today, _ := time.Parse(layout, now)
	ty, tm, td := today.Date()
	by, bm, bd := birthday.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	birthday = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthday) {
		return 0
	}
	age := ty - by
	anniversary := birthday.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age

}

func main() {
	var dob string
	var age int
	fmt.Println("Enter dob in MM-DD-YYYY format")
	fmt.Scan(&dob)
	age = getAge(dob)
	fmt.Println("age=", age)
}
