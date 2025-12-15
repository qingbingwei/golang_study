package condition_sentence

import (
	"fmt"
	"time"
)

func DemoConditionSentence() {
	a := 10
	b := 20

	// if statement
	if a < b {
		fmt.Printf("%d is less than %d\n", a, b)
	} else {
		fmt.Printf("%d is not less than %d\n", a, b)
	}

	// switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// select statement
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Message from channel 1"
	}()

	go func() {
		ch2 <- "Message from channel 2"
	}()

	time.Sleep(1 * time.Second)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
