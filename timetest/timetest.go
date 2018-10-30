package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	timeFormat := "2006-01-02 15:04:05"
	t1, err := time.Parse(timeFormat, "2014-06-15 08:37:18")
	if err == nil {
		fmt.Println("parse success:", t1)
	}
	fmt.Println(t1, err)

	t10 := time.Now().Unix()
	fmt.Println(t10)
	t11 := time.Unix(t10, 0)
	fmt.Println(t10, t11)
}

func Ttest2() {
	//time Add、Sub
	start := time.Now()
	start.Date()

	fmt.Println("start.....")
	end := time.Now()
	d := end.Sub(start)
	fmt.Println(d)
	fmt.Println(d.Hours(), d.Minutes(), d.Seconds())
	d1, _ := time.ParseDuration("1h")
	end2 := start.Add(d1)
	fmt.Println(end2)

	start.Add(-24 * time.Hour)
}
