package main

import (
	"fmt"
	"npsparkbrowser/parks"
)

func main() {
	fmt.Println("Hello world!")

	parks.Configure()
	parks.GetAll()
}
