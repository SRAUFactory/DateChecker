package main

import "flag"
import "fmt"
import "time"

var (
	days = flag.Int("d", 100, "算出する日数")
)

func main() {
	flag.Parse()
	t := time.Now()
	t = t.AddDate(0, 0, *days)
	fmt.Println(t)
}
