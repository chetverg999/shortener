package main

import "fmt"

func main() {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := run(); err != nil {
		panic(err)
	}
}
