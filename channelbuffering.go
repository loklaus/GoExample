package main

import "fmt"

func main()  {
	messages := make(chan string, 2)

	messages <- "hello"
	messages <- "world"
	//messages <- "world"

	go func(){
		fmt.Println(<-messages)
	}()

	fmt.Println(<-messages)
	//fmt.Println(<-messages)
}
