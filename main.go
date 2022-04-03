package main

import (
	"Banking_System/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app.Start()
}
