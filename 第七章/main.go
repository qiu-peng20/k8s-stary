package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		str := os.Getenv("INTERVAL")
		fmt.Println("INTERVAL：", str)
		time.Sleep(2 * time.Second)
	}
}
