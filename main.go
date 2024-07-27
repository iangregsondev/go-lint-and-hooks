package main

import "fmt"

// const xyz = "hello"

func main() {
	// const ggg = "hello"

	_, err := test()
	fmt.Println(err)
	_, err = test()
	if err != nil {
		fmt.Println("error")
		// help
	}

	fmt.Printf("hello")
}

func test() (string, error) {
	return "hello", fmt.Errorf("soem error")
}
