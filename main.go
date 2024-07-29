package main

import "fmt"

//goland:noinspection GoUnusedConst
const xyz = "hello"

const SOMETHING_THAT_IS_NOT_RIGHT = "hello"

//goland:noinspection GoUnusedConst
func main() {
	hello()
	const ggg = "hello"

	myMap := map[string]string{"first key": "first value", "second key": "second value", "third key": "third value", "fourth key": "fourth value", "fifth key": "fifth value"}

	_ = myMap

	fmt.Printf(SOMETHING_THAT_IS_NOT_RIGHT)

	_, err := test()

	fmt.Println("", err)

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
