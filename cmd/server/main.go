package main

func main() {
	if err := run(); err != nil { // Если сервер не запускается по какой-то из причин, программа завершается
		panic(err)
	}
}
