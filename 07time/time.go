package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%.2d.%02d.%04d\n", t.Day(), t.Month(), t.Year())
	//18.05.2019
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	//calculating times
	week := 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)
	//formatting times
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("03 Jan 2006 15:04"))
	s := t.Format("20060102 15:04:05")
	fmt.Println(t, "->", s)

}
