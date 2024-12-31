package main

import "github.com/romanmufid16/go-mongo-redis/app"

func main() {
	server := app.Server()

	err := server.Listen(":3000")
	if err != nil {
		panic("Failed to run server")
	}
}
