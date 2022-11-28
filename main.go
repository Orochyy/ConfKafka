package main

import (
	"ConfKafka/kaffka"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Okay...")

	kaffka.StartKafka()
	fmt.Println("Kaffka started")
	time.Sleep(10 * time.Minute)

}
