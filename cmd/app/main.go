package main

import (
	"fmt"
	"github.com/chetverg999/shortener.git/internal/app"
)

func main() {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
